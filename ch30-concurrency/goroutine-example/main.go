package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("launch goroutine")
	go printNumber()

	fmt.Println("launch goroutine")
	go printNumber()
	time.Sleep(1 * time.Minute)
}

func printNumber() {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}
}
