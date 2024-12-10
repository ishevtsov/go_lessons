package main

import "fmt"

type Funky func(string) string

var f Funky

func main() {
	f = func(s string) string {
		fmt.Printf("Funky %s\n", s)
		return s + s + s
	}
	res := f("Groovy")
	fmt.Println(res)
	fmt.Println("the end!")
}
