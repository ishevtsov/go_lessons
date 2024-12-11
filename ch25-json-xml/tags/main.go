package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID   uint64 `json:"id"`
	Name string `json:"name,omitempty"`
}

func main() {
	p := Product{ID: 42}
	b, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
