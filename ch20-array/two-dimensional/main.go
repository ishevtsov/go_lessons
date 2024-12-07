package main

import (
	"arrays/guestList"
	"fmt"
)

func main() {
	guests := guestList.Generate()
	for _, guest := range guests {
		fmt.Printf("%s %s\n", guest[0], guest[1])
	}
}
