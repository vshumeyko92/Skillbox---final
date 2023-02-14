package billingData

import (
	"Skillbox-diploma/pkg/utils"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

type BillingService struct {
	Data BillingData
}

type BillingServiceInterface interface {
	ReadFile(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() BillingData
	Execute(string) BillingData
}

func GetBillingService() BillingServiceInterface {
	return &BillingService{}
}

// Execute - функция конечной точки сбора данных и возврата связанных данных
func (b *BillingService) Execute(filename string) (result BillingData) {
	path := utils.GetConfigPath(filename)
	bytes, err := b.ReadFile(path)
	if err != nil {
		log.Fatalln("нет данных для билинга: ", err)
	}
	err = b.SetData(bytes)
	if err != nil {
		log.Fatalln(err)
	}
	result = b.DisplayData()
	return
}

// SetData - добавляет данные
func (b *BillingService) SetData(bytes []byte) error {
	intMaskValue, _ := strconv.Atoi(string(bytes))
	if intMaskValue > 255 {
		intMaskValue = intMaskValue / 255
	}

	flagValues := make([]bool, 6)
	sliceIndex := 0
	for i := 1; i <= 255; {
		if sliceIndex+1 > len(flagValues) {
			break
		}
		flagValues[sliceIndex] = intMaskValue&i > 0
		i = i << 1
		sliceIndex++
	}
	b.Data = BillingData{
		CreateCustomer: flagValues[0],
		Purchase:       flagValues[1],
		Payout:         flagValues[2],
		Recurring:      flagValues[3],
		FraudControl:   flagValues[4],
		CheckoutPage:   flagValues[5],
	}
	return nil
}

// DisplayData - показывает инфо по билингу
func (b *BillingService) DisplayData() BillingData {
	return b.Data
}

func (b *BillingService) ReadFile(path string) (res []byte, err error) {
	if len(path) == 0 {
		err = errors.New("такого пути нет")
	}
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer utils.FileClose(file)

	res, err = io.ReadAll(file)
	if err != nil {
		return
	}
	return
}
