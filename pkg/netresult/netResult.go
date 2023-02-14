package netresult

import (
	"Skillbox-diploma/pkg/billingData"
	"Skillbox-diploma/pkg/emailData"
	"Skillbox-diploma/pkg/incidentData"
	"Skillbox-diploma/pkg/mmsData"
	"Skillbox-diploma/pkg/smsData"
	"Skillbox-diploma/pkg/supportData"
	"Skillbox-diploma/pkg/voicecallData"
	"fmt"
	"log"
	"net/http"
)

type ResultSetT struct {
	SMS       [][]smsData.SMSData                `json:"sms"`
	MMS       [][]mmsData.MMSData                `json:"MMS"`
	VoiceCall []voicecallData.VoiceCallData      `json:"voice_call"`
	Email     map[string][][]emailData.EmailData `json:"email"`
	Billing   billingData.BillingData            `json:"billing"`
	Support   []int                              `json:"support"`
	Incident  []incidentData.IncidentData        `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		log.Fatalln(err)
	}
}

// PickDataConnection - функция, которая собирает все функции и выдет единым результатом с помощью родительских структур
func PickDataConnection(w http.ResponseWriter, r *http.Request) {
	sms := smsData.GetSMSService()
	mms := mmsData.GetMMSService()
	v := voicecallData.GetVoiceCallService()
	b := billingData.GetBillingService()
	e := emailData.GetEmailService()
	s := supportData.GetSupportService()
	in := incidentData.GettIncidentService()

	sms.Execute("sms.csv")
	mms.Execute("http://127.0.0.1:8383")
	v.Execute("voicecall.csv")
	b.Execute("billing.cfg")
	e.Execute("email.csv")
	s.Execute("http://127.0.0.1:8484")
	in.Execute("http://127.0.0.1:8585")

	resultSet := ResultSetT{
		SMS:       sms.ReturnFormattedData(),
		MMS:       mms.ReturnFormattedData(),
		VoiceCall: v.ReturnData(),
		Email:     e.ReturnFormattedData(),
		Billing:   b.DisplayData(),
		Support:   s.ReturnFormattedData(),
		Incident:  in.ReturnFormattedData(),
	}
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%v", resultSet)
	if err != nil {
		log.Fatalln(err)
	}
}
