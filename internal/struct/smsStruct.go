package _struct

type SmsData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

var SmsOperator = []string{"Topolo", "Kildy", "Rond"}
