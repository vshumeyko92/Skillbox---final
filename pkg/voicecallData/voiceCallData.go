package voicecallData

import (
	"Skillbox-diploma/pkg/utils"
	"Skillbox-diploma/pkg/validators"
	"errors"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
	"unicode"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           int     `json:"bandwidth"'`
	ResponseTime        int     `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoiceClarity        int     `json:"voice_clarity"`
	MedianCallTime      int     `json:"median_call_time"`
}

type VoiceCallInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() []VoiceCallData
	Execute(string) []VoiceCallData
}

// Execute - функция конечной точки сбора данных и возврата связанных данных
func (c *VoiceCallService) Execute(filename string) []VoiceCallData {
	path := utils.GetConfigPath(filename)
	bytes, err := c.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("нет данных", err)
	}
	err = c.SetData(bytes)
	if err != nil {
		log.Fatalln("нет данных")
	}
	return c.ReturnData()
}

// VoiceCallService - сервис для извлечения и хранения данных о состоянии системы звонков
type VoiceCallService struct {
	Data []VoiceCallData
}

func GetVoiceCallService() VoiceCallInterface {
	return &VoiceCallService{Data: make([]VoiceCallData, 0)}
}

// ReadCSVFile - считываем CSV для получения данных о звонках
func (c *VoiceCallService) ReadCSVFile(path string) (res []byte, err error) {
	if len(path) == 0 {
		err = errors.New("Нет такого пути")
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

// SetData - Добавить данные
func (c *VoiceCallService) SetData(bytes []byte) error {
	initialSize := len(c.Data)
	data := string(bytes[:])
	record := strings.Split(data, "\n")
	for _, rec := range record {
		validated, err := c.ValidateData(rec)
		if err != nil {
			continue
		}
		c.Data = append(c.Data, validated)
	}
	if initialSize == len(c.Data) {
		return errors.New("новых данных нет")
	}
	return nil
}

// ReturnData - показываенные данные звонков
func (c *VoiceCallService) ReturnData() []VoiceCallData {
	return c.Data
}

// ValidateData - валидирует данные
func (c *VoiceCallService) ValidateData(rec string) (validatedData VoiceCallData, err error) {
	trimString := strings.TrimRightFunc(rec, func(r rune) bool {
		return !unicode.IsNumber(r) && !unicode.IsLetter(r)
	})
	attributes := strings.Split(trimString, ";")
	if len(attributes) != reflect.TypeOf(VoiceCallData{}).NumField() {
		err = errors.New("Кол-во параметров не верно")
		return
	}

	country, err := validators.ValidateCountry(attributes[0])
	if err != nil {
		return
	}

	bandwidth, err := validators.ValidateBandwidthInt(attributes[1])
	if err != nil {
		return
	}

	responseTime, err := validators.ValidateInt(attributes[2])
	if err != nil {
		return
	}

	provider, err := validators.ValidateProviderVoiceCall(attributes[3])
	if err != nil {
		return
	}

	connectionStability, err := validators.ValidateConnectionStability(attributes[4])
	if err != nil {
		return
	}

	ttfb, err := validators.ValidateInt(attributes[5])
	if err != nil {
		return
	}

	voiceClarity, err := validators.ValidateInt(attributes[6])
	if err != nil {
		return
	}

	medianCallTime, err := validators.ValidateInt(attributes[7])
	if err != nil {
		return
	}

	validatedData = VoiceCallData{
		Country:             country,
		Bandwidth:           bandwidth,
		ResponseTime:        responseTime,
		Provider:            provider,
		ConnectionStability: connectionStability,
		TTFB:                ttfb,
		VoiceClarity:        voiceClarity,
		MedianCallTime:      medianCallTime,
	}
	return
}
