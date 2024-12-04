package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var agePaul, ageJohn = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "yers old")
	fmt.Println("Paul is", agePaul, "yers old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		if agePaul < ageJohn {
			fmt.Println("Paul is younger than John")
		} else {
			fmt.Println("Paul and John have the same age")
		}
	}
}
