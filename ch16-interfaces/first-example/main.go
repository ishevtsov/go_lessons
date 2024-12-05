package main

import "fmt"

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}

type DomesticAnimal interface {
	ReceiveAffection(from Human)
	GiveAffection(to Human)
	Noiser
}

type Noiser interface {
	Noise()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) ReceiveAffection(from Human) {
	fmt.Printf("The cat named %s has received affection from Human named %s\n", c.Name, from.Firstname)
}

func (c Cat) GiveAffection(to Human) {
	fmt.Printf("The cat named %s has given affection to Human named %s\n", c.Name, to.Firstname)
}

func (c Cat) Noise() {
	fmt.Println("Meooow")
}

func (d Dog) ReceiveAffection(from Human) {
	fmt.Printf("The dog named %s has received affection from Human named %s\n", d.Name, from.Firstname)
}

func (d Dog) GiveAffection(to Human) {
	fmt.Printf("The dog named %s has given affection to Human named %s\n", d.Name, to.Firstname)
}

func (d Dog) Noise() {
	fmt.Println("Whoow")
}

func Pet(animal DomesticAnimal, human Human) {
	animal.GiveAffection(human)
	animal.ReceiveAffection(human)
	animal.Noise()
}

func main() {
	var john Human
	john.Firstname = "John"

	var c Cat
	c.Name = "Maru"

	var d Dog
	d.Name = "Milor"

	Pet(c, john)
	Pet(d, john)
}
