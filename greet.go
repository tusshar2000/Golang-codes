package main

import (
	"fmt"
	"time"
)

func main() {
	timen := time.Now()
	hours := timen.Hour()
	fmt.Println(hours)

	if hours >= 6 && hours < 12 {
		fmt.Println("Good Morning.")
	} else if hours >= 12 && hours < 17 {
		fmt.Println("Good Afternoon.")
	} else if hours >= 17 && hours < 20 {
		fmt.Println("Good Evening.")
	} else {
		fmt.Println("Good Night.")
	}

}
