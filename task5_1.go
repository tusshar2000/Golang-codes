package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int8
	address   []string
}

func defaultName(firstName string, lastName string, age int8, address []string) *person {
	var p person
	p.firstName = firstName
	p.lastName = lastName
	p.age = age
	p.address = address
	return &p
}

func main() {
	person2 := defaultName("Tusshar", "Laddha", 20, []string{"abc", "xyz"})
	person3 := defaultName("Tuss", "Lad", 21, []string{"ac", "yz", "bx"})
	fmt.Println(person2.address)
	fmt.Println(person3.address)

	person1 := person{firstName: "Tusshar",
		lastName: "Laddha",
		age:      20,
		address: []string{"xyz building",
			"abc road",
			"Mumbai"},
	}
	fmt.Println(person1.address)
}
