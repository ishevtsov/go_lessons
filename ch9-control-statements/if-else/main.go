package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var ageJohn, agePaul = rand.Intn(110), rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")
	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		fmt.Println("Paul is younger than John, or both have the same age")
	}
}
