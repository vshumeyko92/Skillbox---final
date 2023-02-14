package mmsData

import (
	"Skillbox-diploma/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type MMSServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []MMSData
	Execute(string) []MMSData
	ReturnFormattedData() [][]MMSData
}

// ReturnFormattedData - отсортированные данные в нужном формате
// Два отсортированных списка
func (s *MMSService) ReturnFormattedData() [][]MMSData {
	fullCountryData := s.displayFullCountry()
	result := make([][]MMSData, 0)
	sort.Sort(ByProviderAsc(fullCountryData))
	result = append(result, fullCountryData)
	sort.Sort(ByCountryAsc(fullCountryData))
	result = append(result, fullCountryData)
	return result
}

// displayFullCountry - Показывает название страны вместо кода alpha-2
func (s *MMSService) displayFullCountry() []MMSData {
	result := s.Data
	countriesMap := utils.GetCountries()
	for i, smsRecord := range s.Data {
		result[i].Country = countriesMap[smsRecord.Country]
	}
	return result
}

type MMSService struct {
	Data []MMSData
}

type ByProviderAsc []MMSData

func (a ByProviderAsc) Len() int {
	return len(a)
}
func (a ByProviderAsc) Less(i, j int) bool {
	return a[i].Provider < a[j].Provider
}

func (a ByProviderAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByCountryAsc []MMSData

func (a ByCountryAsc) Len() int {
	return len(a)
}
func (a ByCountryAsc) Less(i, j int) bool {
	return a[i].Country < a[j].Country
}

func (a ByCountryAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// SendRequest - функция выдает запрос GET
func (m *MMSService) SendRequest(path string) (result []byte, err error) {
	response, err := http.Get(path)
	if err != nil {
		return
	}
	defer utils.CloseReader(response.Body)
	if response.StatusCode != 200 {
		err = fmt.Errorf("Сервер в %s выдает %d: %v", path, response.StatusCode, response.Status)
		return
	}
	result, err = io.ReadAll(response.Body)
	return
}

// DisplayData - показать данные MMS
func (m *MMSService) DisplayData() []MMSData {
	return m.Data
}

// ValidateMMSData - вернуть проверенные данные MMS
func ValidateMMSData(items []MMSData) (validItems []MMSData) {
	for _, v := range items {
		if utils.IsInList(v.Provider, utils.GetProviders()) && utils.IsInList(v.Country, utils.GetCountries()) {
			validItems = append(validItems, v)
		}
	}
	return
}

// SetData - проверка и заполнеение данных в хранилище
func (m *MMSService) SetData(bytes []byte) error {
	initialSize := len(m.Data)

	var newRawData []MMSData
	err := json.Unmarshal(bytes, &newRawData)
	if err != nil {
		return err
	}

	m.Data = append(m.Data, ValidateMMSData(newRawData)...)
	if initialSize == len(m.Data) {
		err := errors.New("Нет новых данных")
		return err
	}
	return nil
}

// Execute- конечная точка для сбора и возврата данных
func (m *MMSService) Execute(s string) []MMSData {
	resp, err := m.SendRequest(s)
	if err != nil {
		log.Fatalln(err)
	}
	err = m.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return m.DisplayData()
}

func GetMMSService() MMSServiceInterface {
	return &MMSService{Data: make([]MMSData, 0)}
}
