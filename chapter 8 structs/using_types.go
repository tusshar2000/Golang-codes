package main

import (
	"fmt"
)

type car struct {
	name     string
	topSpeed float64
}

func main() {
	var Buggati car
	Buggati.name = "Buggati Chiron"
	Buggati.topSpeed = 300
	fmt.Println("Name:", Buggati.name)
	fmt.Println("Top Speed:", Buggati.topSpeed)
}
