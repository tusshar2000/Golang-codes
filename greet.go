package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	currentHour := currentTime.Hour()
	//fmt.Println(hours)

	if currentHour >= 6 && currentHour < 12 {
		fmt.Println("Good Morning.")
	} else if currentHour >= 12 && currentHour < 17 {
		fmt.Println("Good Afternoon.")
	} else if currentHour >= 17 && currentHour < 20 {
		fmt.Println("Good Evening.")
	} else {
		fmt.Println("Good Night.")
	}

}
