package repo

import (
	"Skillbox-diploma/internal/struct"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func BillingReadFile(filePath string) (_struct.BillingData, error) {
	var response _struct.BillingData
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("невозможно прочитать файл: %v", err)
		return response, err
	}
	decNumber := number(string(body))
	binNumber := bitMask(decNumber)
	response = NewBilling(binNumber)
	return response, nil
}

func number(line string) (result uint8) {
	var buffer float64
	for i := 0; i < len(line); i++ {
		index := len(line) - 1 - i
		bit := string(line[index])
		if bit == "1" {
			buffer = buffer + math.Pow(2, float64(i))
		}
	}
	return uint8(buffer)
}

func bitMask(number uint8) (mask string) {
	return strconv.FormatInt(int64(number), 2)
}

func NewBilling(numberBin string) (response _struct.BillingData) {
	length := len(numberBin)
	numberBin = strings.Repeat("0", 6-length) + numberBin
	boolMask := make([]bool, 6)
	for i := 0; i < 6; i++ {
		if numberBin[i:i+1] == "1" {
			boolMask[i] = true
		}
	}
	response.CreateCustomer = boolMask[5]
	response.Purchase = boolMask[4]
	response.Payout = boolMask[3]
	response.Recurring = boolMask[2]
	response.FraudControl = boolMask[1]
	response.CheckoutPage = boolMask[0]
	return
}
