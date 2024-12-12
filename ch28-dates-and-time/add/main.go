package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// + 12 years
	// + 1 Month
	// + 3 days
	now.AddDate(12, 1, 3)
	fmt.Println(now)

	now = now.Add(time.Nanosecond * 1)
	fmt.Println(now)
	now = now.Add(time.Microsecond * 5)
	fmt.Println(now)
	now = now.Add(time.Millisecond * 5)
	fmt.Println(now)
	now = now.Add(time.Second * 5)
	fmt.Println(now)
	now = now.Add(time.Minute * 5)
	fmt.Println(now)
	now = now.Add(time.Hour * 5)
	fmt.Println(now)
}
