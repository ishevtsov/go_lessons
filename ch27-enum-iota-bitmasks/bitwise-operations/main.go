package main

import "fmt"

func main() {
	var x, y, z uint8

	x = 1 // 00000001
	y = 2 // 00000010

	z = x & y
	fmt.Printf("%08b\n", z)
	fmt.Printf("%d\n", z)
}
