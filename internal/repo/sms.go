package repo

import (
	"Skillbox-diploma/internal/struct"
	"encoding/csv"
	"log"
	"os"
)

func SmsReadCsvFile(filePath string, countries map[string]string) (response []_struct.SmsData, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Нет возможности прочитать CSV (sms)"+filePath, err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var buffer _struct.SmsData
	for {
		line, _ := csvReader.Read()
		if line != nil {
			if smsChecker(line, countries) {
				buffer.Country = line[0]
				buffer.Bandwidth = line[1]
				buffer.ResponseTime = line[2]
				buffer.Provider = line[3]
				response = append(response, buffer)
			}
		} else {
			break
		}
	}
	if err != nil {
		log.Printf("нелзя заполнить CSV файл "+filePath, err)
		return
	}
	return response, nil
}

func NewSMS(smsStore *[]_struct.SmsData, filePath string) error {
	recordsToWrite := make([][]string, 0)
	for i := 0; i < len(*smsStore); i++ {
		f0 := (*smsStore)[i].Country
		f1 := (*smsStore)[i].Bandwidth
		f2 := (*smsStore)[i].ResponseTime
		f3 := (*smsStore)[i].Provider
		f := []string{f0, f1, f2, f3}
		recordsToWrite = append(recordsToWrite, f)
	}
	f, err := os.Create(filePath)
	if err != nil {
		log.Printf("невозможно записать файл "+filePath, err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(recordsToWrite)
	if err := w.Error(); err != nil {
		log.Printf("ошибка при записи в csv (sms):", err)
		return err
	}
	return nil
}

func smsChecker(line []string, countries map[string]string) bool {
	if len(line) != 4 {
		return false
	}
	if countries[line[0]] == "" {
		return false
	}
	for i := 0; i < len(_struct.SmsOperator); i++ {
		if line[3] == _struct.SmsOperator[i] {
			return true
		}
	}
	return false
}
