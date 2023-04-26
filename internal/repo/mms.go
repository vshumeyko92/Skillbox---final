package repo

import (
	"Skillbox-diploma/internal/struct"
	"encoding/json"
	"log"
	"net/http"
)

func ReadMMS(httpResponse *http.Response, countries map[string]string) ([]_struct.MMSData, error) {
	var rowData, parsedData []_struct.MMSData
	parsedData = make([]_struct.MMSData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Printf("невозможно разобрать JSON c mms", err)
		return parsedData, err
	}
	for i := 0; i < len(rowData); i++ {
		if mmsCheck(rowData[i], countries) == true {
			parsedData = append(parsedData, rowData[i])
		}
	}
	return parsedData, nil
}

func mmsCheck(line _struct.MMSData, countries map[string]string) bool {
	if countries[line.Country] == "" {
		return false
	}
	for i := 0; i < len(_struct.MmsOperators); i++ {
		if line.Provider == _struct.MmsOperators[i] {
			return true
		}
	}
	return false
}
