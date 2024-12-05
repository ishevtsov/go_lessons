package main

import "fmt"

type Human struct {
	Firstname string
	Lastname  string
	Age       int
	Country   string
}

func (h Human) String() string {
	return fmt.Sprintf("human named %s %s of age %d living in %s", h.Firstname, h.Lastname, h.Age, h.Country)
}

func main() {
	var john Human
	john.Firstname = "John"
	john.Lastname = "Doe"
	john.Age = 45
	john.Country = "USA"

	fmt.Println(john)
}
