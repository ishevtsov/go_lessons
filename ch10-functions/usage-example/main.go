package main

import "fmt"

func main() {
	johnPrice := computePrice(145.80, 3)
	paulPrice := computePrice(26.32, 10)
	robPrice := computePrice(189.21, 2)

	total := johnPrice + paulPrice + robPrice
	fmt.Printf("TOTAL: %0.2f£\n", total)
}

func computePrice(rate float32, nights int) (price float32) {
	price = rate * float32(nights)
	return
}
