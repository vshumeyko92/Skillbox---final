package repo

import (
	"Skillbox-diploma/internal/struct"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func MailReadCsvFile(filePath string, countries map[string]string) ([]_struct.EmailData, error) {
	//Read source file
	var response []_struct.EmailData
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("невозможно прочитать файл "+filePath, err)
		return response, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var buffer _struct.EmailData
	for {
		line, _ := csvReader.Read()
		if line != nil {
			if mailChecker(line, countries) {
				buffer.Country = line[0]
				buffer.Provider = line[1]
				buffer.DeliveryTime, _ = strconv.Atoi(line[2])
				response = append(response, buffer)
			}
		} else {
			break
		}
	}
	if err != nil {
		log.Printf("невозможно внести данные в CSV (email)"+filePath, err)
		return response, err
	}
	return response, nil
}

func NewMail(mailStore *[]_struct.EmailData, filePath string) error {
	recordsToWrite := make([][]string, 0)
	for i := 0; i < len(*mailStore); i++ {
		f0 := (*mailStore)[i].Country
		f1 := (*mailStore)[i].Provider
		f2 := strconv.Itoa((*mailStore)[i].DeliveryTime)
		f := []string{f0, f1, f2}
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
		log.Printf("ошибка при записи в  csv:", err)
		return err
	}
	return nil
}

// mailChecker проверка синтакиса
func mailChecker(line []string, countries map[string]string) bool {
	if len(line) != 3 {
		return false
	}
	if countries[line[0]] == "" {
		return false
	}
	for i := 0; i < len(_struct.EmailOperators); i++ {
		if line[1] == _struct.EmailOperators[i] {
			return true
		}
	}
	return false
}
