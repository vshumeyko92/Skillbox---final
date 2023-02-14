package main

import (
	"Skillbox-diploma/pkg/emailData"
	"fmt"
)

func main() {
	emailService := emailData.GetEmailService()
	fmt.Println(emailService.Execute("email.csv"))
}
