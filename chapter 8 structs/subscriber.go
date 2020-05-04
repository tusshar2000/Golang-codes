package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var subscriber1 subscriber
	subscriber1.name = "Tusshar"
	subscriber1.rate = 500
	subscriber1.active = true

	var subscriber2 subscriber
	subscriber2.name = "Travis"
	subscriber2.rate = 100
	subscriber2.active = false

	fmt.Println(subscriber1.name)
	fmt.Println(subscriber2.name)
}
