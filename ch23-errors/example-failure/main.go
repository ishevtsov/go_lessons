package main

import (
	"errors/error-example/config"
	"fmt"
	"log"
)

func main() {
	confData, err := config.Load()
	if err != nil {
		log.Fatalf("Can't load application config because: %s", err)
	}
	fmt.Println(confData)
}
