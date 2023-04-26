package repo

import (
	"Skillbox-diploma/internal/struct"
	"encoding/json"
	"log"
	"net/http"
)

func ReadSupport(httpResponse *http.Response) ([]_struct.SupportData, error) {
	var rowData, parsedData []_struct.SupportData
	parsedData = make([]_struct.SupportData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Printf("невозможно заполнить JSON Support ", err)
		return parsedData, err
	}
	for i := 0; i < len(rowData); i++ {
		parsedData = append(parsedData, rowData[i])
	}
	return parsedData, nil
}
