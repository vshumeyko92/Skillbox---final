package incidentData

import (
	"Skillbox-diploma/pkg/utils"
	"Skillbox-diploma/pkg/validators"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

type IncidentService struct {
	Data []IncidentData
}

type IncidentServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []IncidentData
	Execute(string) []IncidentData
	ReturnFormattedData() []IncidentData
}

// необходима сортировка по статусам (11.16)
type ByStatus []IncidentData

func (in ByStatus) Len() int {
	return len(in)
}
func (in ByStatus) Less(i, j int) bool {
	return in[i].Status < in[j].Status
}
func (in ByStatus) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

// ReturnFormattedData - возвращает данные об инцидентах
func (is *IncidentService) ReturnFormattedData() []IncidentData {
	result := is.Data
	sort.Sort(ByStatus(result))
	return result
}

// SendRequest - создаём запрос GET
func (is *IncidentService) SendRequest(path string) (result []byte, err error) {
	response, err := http.Get(path)
	if err != nil {
		return
	}
	defer utils.CloseReader(response.Body)
	if response.StatusCode != 200 {
		err = fmt.Errorf("сервер %s выдаёт %d: %v", path, response.StatusCode, response.Status)
		return
	}
	result, err = io.ReadAll(response.Body)
	return

}

// SetData - валидирует и заполняет данные в хранилище
func (is *IncidentService) SetData(bytes []byte) error {
	initialSize := len(is.Data)
	var newRawData []IncidentData
	err := json.Unmarshal(bytes, &newRawData)
	if err != nil {
		return err
	}
	is.Data = append(is.Data, ValidateIncidentData(newRawData)...)
	if initialSize == len(is.Data) {
		err := errors.New("нет новых данных!")
		return err
	}
	return nil
}

// DisplayData - показывает данные инцидентов
func (is *IncidentService) DisplayData() []IncidentData {
	return is.Data
}

// Execute - конечная функция
func (is *IncidentService) Execute(path string) []IncidentData {
	resp, err := is.SendRequest(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = is.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	err = is.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return is.DisplayData()
}

// GetIncidentService - запускает сервис
func GettIncidentService() IncidentServiceInterface {
	return &IncidentService{Data: make([]IncidentData, 0)}
}

// ValidateIncidentData - возвращает отвалидированные данные
func ValidateIncidentData(items []IncidentData) (validItems []IncidentData) {
	for _, v := range items {
		if validators.ValidateIncidentStatus(v.Status) {
			validItems = append(validItems, v)
		}
	}
	return
}
