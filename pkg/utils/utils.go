package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// FileClose - Безошибочное закрытие файла
func FileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalf("ошибка в %v закрытие: %v\n", f, err)
	}
}

// CloseReader - закрыть reader после получения Body
func CloseReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// GetCountries- Получить все доступные страны
func GetCountries() map[string]string {
	return map[string]string{
		"AL": "Albania",
		"AO": "Angola",
		"AM": "Armenia",
		"AU": "Australia",
		"BY": "Belarus",
		"BF": "Burkina Faso",
		"CA": "Canada",
		"CN": "China",
		"HU": "Hungary",
		"JP": "Japan",
		"PL": "Poland",
		"RU": "Russian Federation",
	}
}

// GetProviders- получить доступных провайдеров SMS и MMS
func GetProviders() map[string]struct{} {
	return map[string]struct{}{"Topolo": {}, "Rond": {}, "Kildy": {}}
}

// GetVoiceCallProviders - получить доступных провайдеров звонков
func GetVoiceCallProviders() map[string]struct{} {
	return map[string]struct{}{"TransparentCalls": {}, "E-Voice": {}, "JustPhone": {}}
}

// GetEmailProviders - получить список Email провайдеров
func GetEmailProviders() map[string]struct{} {
	return map[string]struct{}{
		"Gmail": {}, "Yahoo": {}, "Hotmail": {}, "MSN": {}, "Orange": {}, "Comcast": {}, "AOL": {}, "Live": {}, "RediffMail": {}, "GMX": {}, "Protonmail": {}, "Yandex": {}, "Mail.ru": {},
	}
}

// Keys - выдает ключ с мапы
func Keys[Base string | struct{}](m map[string]Base) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}

// IsInList - находиться ли значение в списке (true/false)
func IsInList[Base string | struct{}](search string, list map[string]Base) bool {
	for _, value := range Keys(list) {
		if search == value {
			return true
		}
	}
	return false
}

// GetConfigPath
func GetConfigPath(filename string) (resultPath string) {
	currentLocation, err := os.Getwd()
	if err != nil {
		log.Fatalf("Ошибка при вызове текущего пути: %s\n", err)
	}
	rootFolder := filepath.Dir(filepath.Dir(currentLocation))
	resultPath = filepath.Join(rootFolder, "conf", filename)
	return
}
