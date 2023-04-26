package repo

import (
	"Skillbox-diploma/internal/struct"
	"fmt"
	"sort"
)

func RefreshStatusPage(config _struct.Config, countries map[string]string) (stPage _struct.ResultT) {
	rawData := NewData(config, countries)

	stPage.Data = prepareData(rawData, countries)
	errors, status := errorHandler(rawData.Error, config)
	stPage.Error = errors
	stPage.Status = status
	return
}

func prepareData(rawStruct _struct.RawStruct, countries map[string]string) (stPageSet _struct.ResultSetT) {

	//SMS section
	sms := rawStruct.SMS
	for key := range sms {
		sms[key].Country = countries[sms[key].Country]
	}
	stPageSet.SMS = make([][]_struct.SmsData, 2)
	stPageSet.SMS[0] = make([]_struct.SmsData, len(sms))
	stPageSet.SMS[1] = make([]_struct.SmsData, len(sms))
	sort.Slice(sms, func(i, j int) (less bool) {
		return sms[i].Country < sms[j].Country
	})
	for i := range sms {
		stPageSet.SMS[0][i] = sms[i]
	}
	sort.Slice(sms, func(i, j int) (less bool) {
		return sms[i].Provider < sms[j].Provider
	})
	for i := range sms {
		stPageSet.SMS[1][i] = sms[i]
	}

	//MMS section
	mms := rawStruct.MMS
	for key := range mms {
		mms[key].Country = countries[mms[key].Country]
	}
	stPageSet.MMS = make([][]_struct.MMSData, 2)
	stPageSet.MMS[0] = make([]_struct.MMSData, len(mms))
	stPageSet.MMS[1] = make([]_struct.MMSData, len(mms))
	sort.Slice(mms, func(i, j int) (less bool) {
		return mms[i].Country < mms[j].Country
	})
	for i := range mms {
		stPageSet.MMS[0][i] = mms[i]
	}
	sort.Slice(mms, func(i, j int) (less bool) {
		return mms[i].Provider < mms[j].Provider
	})
	for i := range mms {
		stPageSet.MMS[1][i] = mms[i]
	}

	//Voice section
	stPageSet.VoiceCall = rawStruct.VoiceData

	//Email section
	email := rawStruct.Email
	emailCountries := make(map[string][]_struct.EmailData)
	sortedEmailCountries := make(map[string][]_struct.EmailData)
	stPageSet.Email = make(map[string][][]_struct.EmailData)
	for i := range email {
		emailCountries[email[i].Country] = append(emailCountries[email[i].Country], email[i])
	}
	for key, value := range emailCountries {
		sort.Slice(value, func(i, j int) (less bool) {
			return value[i].DeliveryTime < value[j].DeliveryTime
		})
		sortedEmailCountries[key] = value
	}
	for key, value := range sortedEmailCountries {
		stPageSet.Email[key] = make([][]_struct.EmailData, 2)
		stPageSet.Email[key][0] = make([]_struct.EmailData, 3)
		stPageSet.Email[key][1] = make([]_struct.EmailData, 3)
		stPageSet.Email[key][0] = value[:3]
		stPageSet.Email[key][1] = value[len(value)-4 : len(value)-1]
	}
	fmt.Println(stPageSet.Email)

	//Billing section
	stPageSet.Billing = rawStruct.Billing

	//Support section
	support := rawStruct.Support
	stPageSet.Support = make([]int, 2)
	ticketsCount := 0
	averageTime := 60 / 18
	for i := range support {
		ticketsCount = ticketsCount + support[i].ActiveTickets
	}
	if ticketsCount < 9 {
		stPageSet.Support[0] = 1
	} else if ticketsCount >= 8 && ticketsCount < 16 {
		stPageSet.Support[0] = 2
	} else {
		stPageSet.Support[0] = 3
	}
	stPageSet.Support[1] = averageTime * ticketsCount

	//Incidents section
	//Получите данные об истории инцидентов.
	//Отсортируйте полученные данные так, чтобы все инциденты со статусом active оказались наверху списка, а остальные ниже. Порядок остальной сортировки не важен.
	stPageSet.Incidents = make([]_struct.IncidentData, 0)
	incident := rawStruct.Incidents
	for i := range incident {
		if incident[i].Status == "active" {
			stPageSet.Incidents = append(stPageSet.Incidents, incident[i])
		}
	}

	for i := range incident {
		if incident[i].Status == "closed" {
			stPageSet.Incidents = append(stPageSet.Incidents, incident[i])
		}
	}

	return
}

func errorHandler(errors []error, config _struct.Config) (errorsString string, status bool) {
	if len(errors) == 0 {
		return "", true
	} else {
		if config.Debug {
			for i := range errors {
				errorsString = errorsString + errors[i].Error() + "\n"
			}
		} else {
			errorsString = "Произошли некоторые ошибки при сборе данных из источников"
		}
		status = false
		return
	}
}
