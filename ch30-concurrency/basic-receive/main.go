package main

import "fmt"

func main() {
	var received int
	ch5 := make(chan int, 2)
	ch5 <- 42
	ch5 <- 41
	received = <-ch5
	fmt.Println(received)
}
