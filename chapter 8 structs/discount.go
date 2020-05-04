package main

import (
	"fmt"
)

type subscribers struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscribers) {
	fmt.Println("Name", s.name)
	fmt.Println("Rate", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultsubscribers(name string) *subscribers {
	var s subscribers
	s.name = name
	s.rate = 500
	s.active = true
	return &s
}

func appyDiscount(s *subscribers) {
	s.rate = 100
}

func main() {
	subscriberscriber1 := defaultsubscribers("Tusshar")
	appyDiscount(subscriberscriber1)
	printInfo(subscriberscriber1)

	subscriberscriber2 := defaultsubscribers("Travis")
	printInfo(subscriberscriber2)
}
