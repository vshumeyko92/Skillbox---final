package main

import (
	"Skillbox-diploma/pkg/billingData"
	"fmt"
)

func main() {
	billingService := billingData.GetBillingService()
	fmt.Println(billingService.Execute("billing.cfg"))
}
