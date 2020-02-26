package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func main() {
	alex := person{
		firstName: "Alex", 
		lastName: "Anderson",
		contactInfo: contactInfo{
			email: "alex@gmail.com",
			zipCode: 94000,
		},
	}

	alex.updateName("Hermes")
	alex.print()
}
