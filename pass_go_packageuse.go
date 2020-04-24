package main

import (
	"fmt"
	"keyboard" //we created this package
	"log"
)

func main() {
	fmt.Print("Enter a number: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}
	fmt.Println(status)

}
