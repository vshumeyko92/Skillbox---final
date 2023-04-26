package main

import (
	"Skillbox-diploma/internal/repo"
	"Skillbox-diploma/internal/web"
)

func main() {
	config := repo.ConfigReader()
	countries := repo.GetCountries()
	serverUrl := config.WebListernerAddress
	web.Router(serverUrl, config, countries)
}
