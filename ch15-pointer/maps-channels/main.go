package main

import "fmt"

func main() {
	cities := make(map[string]string)
	addElement(cities)
	fmt.Println(cities)
}

func addElement(cities map[string]string) {
	cities["France"] = "Paris"
}
