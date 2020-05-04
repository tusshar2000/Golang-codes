package main

import (
	"fmt"
)

type subs struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subs) {
	fmt.Println("Name", s.name)
	fmt.Println("Rate", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultsubs(name string) subs {
	var s subs
	s.name = name
	s.rate = 500
	s.active = true
	return s
}

func main() {
	subs1 := defaultsubs("Tusshar")
	subs1.rate = 100
	printInfo(subs1)
}
