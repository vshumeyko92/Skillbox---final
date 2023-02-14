package validators

import (
	"Skillbox-diploma/pkg/utils"
	"errors"
	"strconv"
)

// ValidateCountry- проверяем, согласно списку, есть ли данные о стране
func ValidateCountry(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetCountries()) {
		result = rawValue
	} else {
		err = errors.New("нет такой страны")
	}
	return
}

// ValidateBandwidth- проверям правильность данных пропускной способности канала
func ValidateBandwidth(rawValue string) (result string, err error) {
	bandwidth, err := strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	if bandwidth < 0 || bandwidth > 100 {
		err = errors.New("процент пропускной способности канала не верна! (вне диапазона от 0 до 100%-)")
		return
	}
	result = rawValue
	return
}

// ValidateResponseTime - проверяем правильность данны о среднем времени ответа
func ValidateResponseTime(rawValue string) (result string, err error) {
	_, err = strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	result = rawValue
	return
}

// ValidateProvider- проверям есть ли провайдер для SMS и MMS
func ValidateProvider(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetProviders()) {
		result = rawValue
	} else {
		err = errors.New("такого провайдера нет")
	}
	return
}

// ValidateProviderVoiceCall - проверяем еть ли провайдер для звонков
func ValidateProviderVoiceCall(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetVoiceCallProviders()) {
		result = rawValue
	} else {
		err = errors.New("такого провайдера нет")
	}
	return
}

// ValidateProviderEmail- проверяем еть ли провайдер для email
func ValidateProviderEmail(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetEmailProviders()) {
		result = rawValue
	} else {
		err = errors.New(" такого провайдера нет")
	}
	return
}

// ValidateBandwidthInt- проверям правильность данных пропускной способности канала (для задания с VoiceCall)
func ValidateBandwidthInt(rawValue string) (result int, err error) {
	bandwidth, err := strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	if bandwidth < 0 || bandwidth > 100 {
		err = errors.New("процент пропускной способности канала не верна! (вне диапазона от 0 до 100%-)")
		return
	}
	result = bandwidth
	return
}

// ValidateInt - проверяем данные на соотвествие целочисленным
func ValidateInt(rawValue string) (result int, err error) {
	_, err = strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	return
}

// ValidateConnectionStability - проверяем данные стабильности соединения (в диапазоне от 0 до 1)
func ValidateConnectionStability(rawValue string) (value float32, err error) {
	value64, err := strconv.ParseFloat(rawValue, 32)
	if err != nil {
		return
	}
	if value64 >= 0 && value64 <= 1 {
		value = float32(value64)
	} else {
		err = errors.New("вне диапазона")
	}
	return
}

// ValidateIncidentStatus - проверяет статус active/closed
func ValidateIncidentStatus(statusValue string) bool {
	return utils.IsInList(statusValue, map[string]struct{}{"active": {}, "closed": {}})
}
