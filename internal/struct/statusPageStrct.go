package _struct

type RawStruct struct {
	SMS       []SmsData
	MMS       []MMSData
	VoiceData []VoiceCallData
	Email     []EmailData
	Billing   BillingData
	Support   []SupportData
	Incidents []IncidentData
	Error     []error
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]SmsData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}
