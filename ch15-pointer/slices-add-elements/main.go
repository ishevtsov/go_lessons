package main

import (
	"fmt"
	"log"
)

func main() {
	EUcoutries := []string{"Austria", "Belgium", "Bulgaria"}
	addCountries(EUcoutries)
	fmt.Println(EUcoutries)
}

func addCountries(countries []string) {
	countries = append(countries, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
	log.Println(countries)
}
