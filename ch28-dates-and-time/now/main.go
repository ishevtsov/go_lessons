package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%s\n", now)
	fmt.Println(now.Unix())
}
