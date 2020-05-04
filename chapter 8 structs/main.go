package main

import (
	"fmt"
	"magazine"
)

func main() {
	s := magazine.Subscribers{Rate: 500} //we can omit some fields
	fmt.Println(s.Rate)
	//struct literals
	subscriber := magazine.Subscribers{Name: "Tusshar"}
	subscriber.Building = "heights"
	subscriber.Street = "Carter Road"
	fmt.Println(subscriber.Name)
	fmt.Println(subscriber.Building)
	fmt.Println(subscriber.Street)
}
