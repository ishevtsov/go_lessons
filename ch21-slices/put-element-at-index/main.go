package main

import "fmt"

func main() {
	b := []string{"john", "jeanne", "jean", "josh"}

	// 1) add an element to the end of the slice
	b = append(b, "")
	fmt.Println(b)

	// 2) copy b[i:] to b[i+1:]
	copy(b[2:], b[1:])
	fmt.Println(b)

	// 3) set "C" at index 2
	b[1] = "joe"
	fmt.Println(b)
}
