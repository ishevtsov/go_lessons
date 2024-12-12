package main

import (
	"fmt"
	"time"
)

func main() {
	start, err := time.Parse("2006-01-02", "2023-02-10")
	if err != nil {
		panic(err)
	}

	end, err := time.Parse("2006-01-02", "2024-12-12")
	if err != nil {
		panic(err)
	}

	for i := start; i.Unix() < end.Unix(); i = i.AddDate(0, 0, 1) {
		fmt.Println(i.Format(time.RFC3339))
	}
}
