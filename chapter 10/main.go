package main

import (
	"calendar"
	"fmt"
	"log"
)

func main() {
	// date := calendar.Date{}
	event := calendar.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Day())
	fmt.Println(event.Month())
	fmt.Println(event.Year()) //fmt.Println(event.Date.Year())
}
