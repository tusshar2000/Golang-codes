package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var number int

	rand.Seed(time.Now().UnixNano()) //time seed used to generate random numbers upon every execution.
	target := rand.Intn(50) + 1      //randomly generated target

	fmt.Println("Guess the target game starts")
	flag := false

	for guess := 0; guess < 5; guess++ {
		if guess == 4 { //5-4==1, used for correct usage of grammar, guess and guesses
			fmt.Println("You have", 5-guess, "guess left!") //when 1 guess is left we need to use guess.
		} else {
			fmt.Println("You have", 5-guess, "guesses left!") //else use guesses.
		}
		fmt.Println("Enter number")

		_, err := fmt.Scanf("%d\n", &number) //user input guess
		if err != nil {
			log.Fatal(err)
		}

		if number < target {
			fmt.Println("Your guess is low")
		} else if number > target {
			fmt.Println("Your guess is high")
		} else {
			flag = true
			fmt.Println("Congrats, you made the correct guess.")
			break
		}
	}
	if flag == false {
		fmt.Println("Sorry you couldn't guess, the number was:", target)
	}

}

// sample output

// C:\Go-Internship\src>go run randomNo_task.go
// Guess the target game starts
// You have 5 guesses left!
// Enter number
// 25
// Your guess is low
// You have 4 guesses left!
// Enter number
// 35
// Your guess is high
// You have 3 guesses left!
// Enter number
// 30
// Your guess is high
// You have 2 guesses left!
// Enter number
// 27
// Your guess is low
// You have 1 guess left!
// Enter number
// 28
// Congrats, you made the correct guess.
