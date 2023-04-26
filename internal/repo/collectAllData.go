package repo

import (
	"Skillbox-diploma/internal/struct"
	"errors"
	"log"
	"net/http"
)

func NewData(config _struct.Config, countries map[string]string) (rawStruct _struct.RawStruct) {
	rawStruct.Error = make([]error, 0)

	smsDataPath := config.SmsSource
	var smsErr error
	rawStruct.SMS, smsErr = SmsReadCsvFile(smsDataPath, countries)
	if smsErr != nil {
		log.Printf("ошибка в чтении данных SMS: %s", smsErr.Error())
		rawStruct.Error = append(rawStruct.Error, smsErr)
	}
	//write slice "sms" to CSV file
	smsDataSave := config.SmsTarget
	if err := NewSMS(&rawStruct.SMS, smsDataSave); err != nil {
		log.Printf("ошибка в записи данных в CSV (SMS): %s", err.Error())
		rawStruct.Error = append(rawStruct.Error, err)
	}

	mmsDataUrl := config.MmsSource
	response, mmsErr := http.Get(mmsDataUrl)
	if mmsErr != nil {
		log.Printf("ошибка web запроса MMS: %s", mmsErr.Error())
		rawStruct.Error = append(rawStruct.Error, mmsErr)
	} else if response.StatusCode != 200 {
		log.Printf("ошибка при web запросе MMS: %s", response.Status)
		rawStruct.Error = append(rawStruct.Error, errors.New("MMS код ответа не равен 200"))
	} else {
		rawStruct.MMS, mmsErr = ReadMMS(response, countries)
	}
	if mmsErr != nil {
		log.Printf("ошибка web запроса MMS: %s", mmsErr.Error())
		rawStruct.Error = append(rawStruct.Error, mmsErr)
	}

	voiceDataPath := config.VoiceSource
	var voiceErr error
	rawStruct.VoiceData, voiceErr = VoiceReadCsvFile(voiceDataPath, countries)
	if voiceErr != nil {
		log.Printf("ошибка чтения данных voice: %s", voiceErr.Error())
		rawStruct.Error = append(rawStruct.Error, voiceErr)
	}

	//write slice "voice" to CSV file
	voiceDataSave := config.VoiceTarget
	if voiceErr := NewVC(&rawStruct.VoiceData, voiceDataSave); voiceErr != nil {
		log.Printf("ошибка при записи в CSV (voice): %s", voiceErr.Error())
		rawStruct.Error = append(rawStruct.Error, voiceErr)
	}

	mailDataPath := config.MailSource
	var mailErr error
	rawStruct.Email, mailErr = MailReadCsvFile(mailDataPath, countries)
	if mailErr != nil {
		log.Printf("ошибка чтения данных EMail: %s", mailErr.Error())
		rawStruct.Error = append(rawStruct.Error, mailErr)
	}
	//write slice "mail" to CSV file
	mailDataSave := config.MailTarget
	if mailErr = NewMail(&rawStruct.Email, mailDataSave); mailErr != nil {
		log.Printf("ошибка записи данных в CSV (email): %s", mailErr.Error())
		rawStruct.Error = append(rawStruct.Error, mailErr)
	}

	var billingErr error
	billingDataPath := config.BillingSource
	rawStruct.Billing, billingErr = BillingReadFile(billingDataPath)
	if billingErr != nil {
		log.Printf("ошибка чтения данных Billing: %s", billingErr.Error())
		rawStruct.Error = append(rawStruct.Error, billingErr)
	}

	supportDataUrl := config.SupportSource
	var supportErr error
	response, supportErr = http.Get(supportDataUrl)
	if supportErr != nil {
		log.Printf("ошибка при web запросе Support : %s", supportErr.Error())
		rawStruct.Error = append(rawStruct.Error, supportErr)
	} else if response.StatusCode != 200 {
		log.Printf("ошибка при web запросе Support: %s", response.Status)
		rawStruct.Error = append(rawStruct.Error, errors.New("код не равен 200 (support)"))
	} else {
		rawStruct.Support, supportErr = ReadSupport(response)
	}
	if supportErr != nil {
		log.Printf("ошибка при web запросе Support: %s", supportErr.Error())
		rawStruct.Error = append(rawStruct.Error, supportErr)
	}

	accendentDataUrl := config.AccendendSource
	var accendentErr error
	response, accendentErr = http.Get(accendentDataUrl)
	if accendentErr != nil {
		log.Printf("ошибка при web запросе accendent: %s", accendentErr.Error())
		rawStruct.Error = append(rawStruct.Error, accendentErr)
	} else if response.StatusCode != 200 {
		log.Printf("ошибка при web запросе accendent: %s", response.Status)
		rawStruct.Error = append(rawStruct.Error, errors.New("Accendent код ответа не равен 200"))
	} else {
		rawStruct.Incidents, accendentErr = ReadAccendent(response)
	}
	if accendentErr != nil {
		log.Printf("ошибка при web запросе accendent: %s", accendentErr.Error())
		rawStruct.Error = append(rawStruct.Error, accendentErr)
	}

	return rawStruct
}
