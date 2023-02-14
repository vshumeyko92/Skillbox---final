package main

import (
	"Skillbox-diploma/pkg/smsData"
	"fmt"
)

func main() {
	smsService := smsData.GetSMSService()
	smsService.Execute("sms.csv")
	fmt.Println(smsService.DisplayData())
}
