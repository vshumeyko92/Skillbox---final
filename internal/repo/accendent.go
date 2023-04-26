package repo

import (
	"Skillbox-diploma/internal/struct"
	"encoding/json"
	"log"
	"net/http"
)

func ReadAccendent(httpResponse *http.Response) ([]_struct.IncidentData, error) {
	var rowData, parsedData []_struct.IncidentData
	parsedData = make([]_struct.IncidentData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Printf("невозможно разобрать JSON c инцидентами ", err)
		return parsedData, err
	}
	for i := 0; i < len(rowData); i++ {
		parsedData = append(parsedData, rowData[i])
	}
	return parsedData, nil
}
