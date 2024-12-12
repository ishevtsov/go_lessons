package main

import (
	"fmt"
	"time"
)

func main() {
	timeToParse := "2024-12-12T07:33:10+00:00"
	t, err := time.Parse("2006-01-02T15:04:05-07:00", timeToParse)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
}
