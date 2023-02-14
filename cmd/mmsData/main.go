package main

import (
	"Skillbox-diploma/pkg/mmsData"
	"fmt"
)

func main() {
	mmsService := mmsData.GetMMSService()
	fmt.Println(mmsService.Execute("http://127.0.0.1:8383"))
}
