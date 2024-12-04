package main

import (
	"fmt"
	"rules2/invoice"
)

func init() {
	fmt.Println("main")
}

func main() {
	fmt.Println("--program start--")
	invoice.Print()
}
