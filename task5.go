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

func main() {
	person2 := person{}
	person2.firstName = "Tusshar"
	person2.lastName = "Laddha"
	person2.age = 20
	// person.address = append(person.address, "abc building", "xyz road", "Mumbai")
	person2.address = make([]string, 3)
	person2.address[0] = "abc building"
	person2.address[1] = "xyz road"
	person2.address[2] = "Mumbai"

	fmt.Println(person2.address)
}
