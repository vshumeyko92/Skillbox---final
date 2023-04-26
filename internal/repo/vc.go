package repo

import (
	"Skillbox-diploma/internal/struct"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func VoiceReadCsvFile(filePath string, countries map[string]string) ([]_struct.VoiceCallData, error) {
	var response []_struct.VoiceCallData
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("невозможно прочитать входящий файл "+filePath, err)
		return response, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var buffer _struct.VoiceCallData
	for {
		line, _ := csvReader.Read()
		if line != nil {
			if VoiceCheck(line, countries) {
				buffer.Country = line[0]
				buffer.Bandwidth, _ = strconv.Atoi(line[1])
				buffer.ResponseTime, _ = strconv.Atoi(line[2])
				buffer.Provider = line[3]
				ConnectionStability, _ := strconv.ParseFloat(line[4], 32)
				buffer.ConnectionStability = float32(ConnectionStability)
				buffer.TTFB, _ = strconv.Atoi(line[5])
				buffer.VoiceClarity, _ = strconv.Atoi(line[6])
				buffer.MedianCallTime, _ = strconv.Atoi(line[7])
				response = append(response, buffer)
			}
		} else {
			break
		}
	}
	if err != nil {
		log.Printf("невозможно заполнить CSV"+filePath, err)
	}
	return response, err
}

func NewVC(voiceStore *[]_struct.VoiceCallData, filePath string) error {
	recordsToWrite := make([][]string, 0)
	for i := 0; i < len(*voiceStore); i++ {
		f0 := (*voiceStore)[i].Country
		f1 := strconv.Itoa((*voiceStore)[i].Bandwidth)
		f2 := strconv.Itoa((*voiceStore)[i].ResponseTime)
		f3 := (*voiceStore)[i].Provider
		f4 := fmt.Sprintf("%f", (*voiceStore)[i].ConnectionStability)
		f5 := strconv.Itoa((*voiceStore)[i].TTFB)
		f6 := strconv.Itoa((*voiceStore)[i].VoiceClarity)
		f7 := strconv.Itoa((*voiceStore)[i].MedianCallTime)
		f := []string{f0, f1, f2, f3, f4, f5, f6, f7}
		recordsToWrite = append(recordsToWrite, f)
	}
	f, err := os.Create(filePath)
	if err != nil {
		log.Printf("невозможно заполнить файл "+filePath, err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(recordsToWrite)
	if err := w.Error(); err != nil {
		log.Printf("ошибка при записи csv:", err)
		return err
	}
	return nil
}

func VoiceCheck(line []string, countries map[string]string) bool {
	//Syntax check, according the rules
	if len(line) != 8 {
		return false
	}
	if countries[line[0]] == "" {
		return false
	}
	for i := 0; i < len(_struct.VoiceOperators); i++ {
		if line[3] == _struct.VoiceOperators[i] {
			return true
		}
	}
	return false
}
