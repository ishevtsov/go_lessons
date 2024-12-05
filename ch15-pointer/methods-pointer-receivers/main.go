package main

import "fmt"

type Cat struct {
	Color string
	Age   uint8
	Name  string
}

func (c *Cat) Meow() {
	fmt.Println("Meooooow")
}

func (c *Cat) Rename(newName string) {
	c.Name = newName
}

func (c Cat) Rename2(newName string) {
	c.Name = newName
}

func main() {
	cat := Cat{Color: "blue", Age: 8, Name: "Milow"}
	cat.Rename("Bob")
	fmt.Println(cat.Name)

	cat.Rename2("Ben")
	fmt.Println(cat.Name)
}
