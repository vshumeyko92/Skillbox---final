package emailData

import (
	"Skillbox-diploma/pkg/utils"
	"Skillbox-diploma/pkg/validators"
	"errors"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type EmailService struct {
	Data []EmailData
}

type EmailServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []EmailData
	Execute(string) map[string][][]EmailData
	ReturnFormattedData() map[string][][]EmailData
}

func GetEmailService() EmailServiceInterface {
	return &EmailService{Data: make([]EmailData, 0)}
}

// Execute - функция конечной точки сбора данных и возврата связанных данных
func (e *EmailService) Execute(filename string) map[string][][]EmailData {
	path := utils.GetConfigPath(filename)
	bytes, err := e.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("нет данных для сервиса email: ", err)
	}
	err = e.SetData(bytes)
	if err != nil {
		log.Fatalln("невозможно использовать данные")
	}
	return e.ReturnFormattedData()
}

type ByMinDelTime []EmailData

func (a ByMinDelTime) Len() int {
	return len(a)
}
func (a ByMinDelTime) Less(i, j int) bool {
	return a[i].DeliveryTime < a[j].DeliveryTime
}
func (a ByMinDelTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByMaxDelTime []EmailData

func (a ByMaxDelTime) Len() int {
	return len(a)
}
func (a ByMaxDelTime) Less(i, j int) bool {
	return a[i].DeliveryTime < a[j].DeliveryTime
}
func (a ByMaxDelTime) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// displayFullCountry - вместо alpha-2 выдаст название страны
func (e *EmailService) displayFullCountry() []EmailData {
	result := e.Data
	countriesMap := utils.GetCountries()
	for i, emailRec := range result {
		result[i].Country = countriesMap[emailRec.Country]
	}
	return result
}

// ReturnFormattedData - форматирует данные
func (e *EmailService) ReturnFormattedData() map[string][][]EmailData {
	result := make(map[string][][]EmailData)

	rawData := make(map[string][]EmailData)

	for _, value := range e.displayFullCountry() {
		rawData[value.Country] = append(rawData[value.Country], value)
	}

	for key, valuesList := range rawData {
		minTimeProviders := make([]EmailData, 0)
		minTimeProviders = append(minTimeProviders, valuesList...)
		sort.Sort(ByMinDelTime(minTimeProviders))

		maxTimeProviders := make([]EmailData, 0)
		maxTimeProviders = append(maxTimeProviders, valuesList...)
		sort.Sort(ByMaxDelTime(maxTimeProviders))

		if len(maxTimeProviders) > 3 {
			maxTimeProviders = maxTimeProviders[:2]
		}

		if len(minTimeProviders) > 3 {
			minTimeProviders = minTimeProviders[:2]
		}

		result[key] = append(result[key], minTimeProviders, maxTimeProviders)
	}
	return result
}

// ReadCSVFile - считывает данные csv файла и выдает данные
func (e *EmailService) ReadCSVFile(path string) (res []byte, err error) {
	if len(path) == 0 {
		err = errors.New("ошибка в пути")
	}

	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer utils.FileClose(file)

	res, err = io.ReadAll(file)
	if err != nil {
		return
	}
	return
}

// SetData - добавляет данные
func (e *EmailService) SetData(bytes []byte) error {
	initialSize := len(e.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := e.ValidateData(record)
		if err != nil {
			continue
		}
		e.Data = append(e.Data, validated)
	}
	if initialSize == len(e.Data) {
		return errors.New("no new data received")
	}
	return nil
}

// DisplayData - показывает данные email
func (e *EmailService) DisplayData() []EmailData {
	return e.Data
}

// ValidateData - получить массив данных из строки (если возможно)
func (e *EmailService) ValidateData(record string) (validatedData EmailData, err error) {
	cleanString := strings.TrimRightFunc(record, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	attrs := strings.Split(cleanString, ";")
	if len(attrs) != reflect.TypeOf(EmailData{}).NumField() {
		err = errors.New("amount of parameters provided is wrong")
		return
	}

	country, err := validators.ValidateCountry(attrs[0])
	if err != nil {
		return
	}

	provider, err := validators.ValidateProviderEmail(attrs[1])
	if err != nil {
		return
	}

	deliveryTime, err := validators.ValidateInt(attrs[2])
	if err != nil {
		return
	}

	validatedData = EmailData{
		Country:      country,
		Provider:     provider,
		DeliveryTime: deliveryTime,
	}

	return
}
