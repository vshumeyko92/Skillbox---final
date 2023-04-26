package _struct

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

var EmailOperators = []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast",
	"AOL", "Live", "RediffMail", "GMX", "Proton Mail", "Yandex", "Mail.ru"}
