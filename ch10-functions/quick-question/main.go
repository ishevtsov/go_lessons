package main

import "fmt"

func main() {
	const HotelName = "Golang"
	printHeader()
}

func printHeader() {
	fmt.Println("Hotel", HotelName)
	fmt.Println("San Francisco, CA")
}
