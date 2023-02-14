package main

import (
	"Skillbox-diploma/pkg/incidentData"
	"fmt"
)

func main() {
	incidentService := incidentData.GettIncidentService()
	fmt.Println(incidentService.Execute("http://127.0.0.1:8585"))
}
