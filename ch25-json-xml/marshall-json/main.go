package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID   uint64   `json:"id"`
	Name string   `json:"name"`
	SKU  string   `json:"sku"`
	Cat  Category `json:"category"`
}

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func main() {
	p := Product{
		ID:   42,
		Name: "Tea Pot",
		SKU:  "TP12",
		Cat: Category{
			ID:   2,
			Name: "Tea",
		}}

	b, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
