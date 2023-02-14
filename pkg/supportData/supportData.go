package supportData

import (
	"Skillbox-diploma/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type SupportService struct {
	Data []SupportData
}

type SupportServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []SupportData
	Execute(string) []SupportData
	ReturnFormattedData() []int
}

// ReturnFormattedData - функция выводит состояние загрузки (1,2,3) (Задание 11.14)
func (s *SupportService) ReturnFormattedData() []int {
	result := make([]int, 0)
	const (
		LowSupportLoad    int = 1
		MediumSupportLoad int = 2
		HighSupportLoad   int = 3
		TicketsPerHour    int = 18
		HighLevelLoad     int = 16
		LowLevelLoad      int = 9
	)
	activeTicketNumber := 0
	for _, ticket := range s.Data {
		activeTicketNumber += ticket.ActiveTickets
	}
	expextedTime := (60 / TicketsPerHour) * activeTicketNumber
	currentLoad := MediumSupportLoad

	if activeTicketNumber < LowLevelLoad {
		currentLoad = LowSupportLoad
	} else {
		if activeTicketNumber > HighLevelLoad {
			currentLoad = HighSupportLoad
		}
	}
	result = append(result, currentLoad, expextedTime)
	return result
}

// DisplayData - показывает данные
func (s *SupportService) DisplayData() []SupportData {
	return s.Data
}

// Execute - функция конечной точки сбора данных и возврата связанных данных
func (s *SupportService) Execute(path string) []SupportData {
	resp, err := s.SendRequest(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = s.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return s.DisplayData()
}

// SendRequest - функция которая делает GET запрпос
func (s *SupportService) SendRequest(path string) (result []byte, err error) {
	response, err := http.Get(path)
	if err != nil {
		return
	}
	defer utils.CloseReader(response.Body)
	if response.StatusCode != 200 {
		err = fmt.Errorf("Сервыер %s выдаёт %d: %v", path, response.StatusCode, response.Status)
		return
	}
	result, err = io.ReadAll(response.Body)
	return
}

// ValidateSupportData - возвращает отвалидированные данные
func ValidateSupportData(items []SupportData) (validItems []SupportData) {
	for _, v := range items {
		if v.ActiveTickets >= 0 {
			validItems = append(validItems, v)
		}
	}
	return
}

// SetData -валидирует данные и отправляет  в хранилище
func (s *SupportService) SetData(bytes []byte) error {
	initialSize := len(s.Data)
	var newRawData []SupportData
	err := json.Unmarshal(bytes, &newRawData)
	if err != nil {
		return err
	}

	s.Data = append(s.Data, ValidateSupportData(newRawData)...)
	if initialSize == len(s.Data) {
		err := errors.New("Нет новых данных")
		return err
	}
	return nil
}

// GetSupportService
func GetSupportService() SupportServiceInterface {
	return &SupportService{Data: make([]SupportData, 0)}
}
