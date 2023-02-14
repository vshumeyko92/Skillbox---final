package main

import (
	"Skillbox-diploma/pkg/supportData"
	"fmt"
)

func main() {
	supportService := supportData.GetSupportService()
	fmt.Println(supportService.Execute("http://127.0.0.1:8484"))
}
