package main

import "fmt"

func main() {
	s := "我爱 Golang"
	for _, v := range s {
		// v is of type rune
		fmt.Printf("Unicode code point: %U - character '%c' - binary %b - hex %X - Decimal %d\n", v, v, v, v, v)
	}
}
