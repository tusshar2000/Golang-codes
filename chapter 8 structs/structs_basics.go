package main

import (
	"fmt"
)

func main() {
	var subscriber struct {
		name   string
		rate   float64
		active bool
	}
	subscriber.name = "Tusshar"
	subscriber.rate = 500
	subscriber.active = true

	fmt.Println(subscriber.rate)
}
