package smsData

import (
	"Skillbox-diploma/pkg/utils"
	"Skillbox-diploma/pkg/validators"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// SMSData - структура хранения данных:
// Country - Alpha-2 код страны
// Bandwidth - пропускная способность канала (от 0 до 100)
// ResponseTime - Среднее время ответа в миллисекундах
// Provider - название компании провайдера
type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type SMSService struct {
	Data []SMSData
}

type SMSServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []SMSData
	Execute(string) []SMSData
	ReturnFormattedData() [][]SMSData
}

// ReturnFormattedData - отсортированные данные в нужном формате
// Два отсортированных списка
func (s *SMSService) ReturnFormattedData() [][]SMSData {
	fullCountryData := s.DisplayFullCountry()
	result := make([][]SMSData, 0)
	sort.Sort(ByProviderAsc(fullCountryData))
	result = append(result, fullCountryData)
	sort.Sort(ByCountryAsc(fullCountryData))
	result = append(result, fullCountryData)
	return result
}

// Execute - функция конечной точки сбора данных и возврата связанных данных
func (s *SMSService) Execute(filename string) []SMSData {
	path := utils.GetConfigPath(filename)
	bytes, err := s.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("нет данных: ", err)
	}
	err = s.SetData(bytes)
	if err != nil {
		log.Fatalln("невозможно использовать данные:", err)
	}
	return s.DisplayData()
}

// GetSMSService - сделать сервис для данных SMS
func GetSMSService() SMSServiceInterface {
	return &SMSService{Data: make([]SMSData, 0)}
}

type ByProviderAsc []SMSData

func (a ByProviderAsc) Len() int {
	return len(a)
}
func (a ByProviderAsc) Less(i, j int) bool {
	return a[i].Provider < a[j].Provider
}

func (a ByProviderAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByCountryAsc []SMSData

func (a ByCountryAsc) Len() int {
	return len(a)
}
func (a ByCountryAsc) Less(i, j int) bool {
	return a[i].Country < a[j].Country
}

func (a ByCountryAsc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// SetData - Добавить данные из содержимого файла
func (s *SMSService) SetData(bytes []byte) error {
	initialSize := len(s.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := s.ValidateData(strings.Trim(record, "\n"))
		if err != nil {
			continue
		}
		s.Data = append(s.Data, validated)
	}
	if initialSize == len(s.Data) {
		return errors.New("нет новых данных")
	}
	return nil
}

// ValidateData - получить массив данных из строки (если возможно)
func (s *SMSService) ValidateData(record string) (validatedData SMSData, err error) {
	cleanString := strings.TrimRightFunc(record, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	attributes := strings.Split(cleanString, ";")
	if len(attributes) != reflect.TypeOf(SMSData{}).NumField() {
		err = errors.New("кол-во параметров- неверное")
		return
	}

	country, err := validators.ValidateCountry(attributes[0])
	if err != nil {
		return
	}

	bandwidth, err := validators.ValidateBandwidth(attributes[1])
	if err != nil {
		return
	}

	responseTime, err := validators.ValidateProvider(attributes[3])
	if err != nil {
		return
	}

	provider, err := validators.ValidateProvider(attributes[3])
	if err != nil {
		return
	}

	validatedData = SMSData{
		Country:      country,
		Bandwidth:    bandwidth,
		ResponseTime: responseTime,
		Provider:     provider,
	}
	return
}

// DisplayFullCountry - Показывает название страны вместо кода alpha-2
func (s *SMSService) DisplayFullCountry() []SMSData {
	result := s.Data
	countriesMap := utils.GetCountries()
	for i, smsRecord := range s.Data {
		result[i].Country = countriesMap[smsRecord.Country]
	}
	return result
}

// DisplayData - показывает данные SMS
func (s *SMSService) DisplayData() []SMSData {
	return s.Data
}

func (s *SMSService) ReadCSVFile(path string) (res []byte, err error) {
	if len(path) == 0 {
		err = errors.New("путь не указан")
	}
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer utils.FileClose(file)
	res, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return
}
