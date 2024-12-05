package main

import "fmt"

func main() {
	EUcountries := []string{"Austria", "Belgium", "Bulgaria"}
	addCountries(&EUcountries)
	fmt.Println(EUcountries)
}

func addCountries(countriesPtr *[]string) {
	*countriesPtr = append(*countriesPtr, []string{"Croatia", "Republic of Cyprus", "Czech Republic", "Denmark", "Estonia", "Finland", "France", "Germany", "Greece", "Hungary", "Ireland", "Italy", "Latvia", "Lithuania", "Luxembourg", "Malta", "Netherlands", "Poland", "Portugal", "Romania", "Slovakia", "Slovenia", "Spain", "Sweden"}...)
}
