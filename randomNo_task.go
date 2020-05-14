package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(50) + 1
}

func guessLeft(guess int) {
	//used for correct usage of grammar, "guess" and "guesses".
	if guess == 4 {
		fmt.Println("You have", 5-guess, "guess left!")
	} else {
		fmt.Println("You have", 5-guess, "guesses left!")
	}
}

func checkGuess(number int, target int, flag *bool) string {
	output := ""
	if number <= 0 || number > 50 {
		*flag = false
		output = "Please enter a number between 1-50"
		return output
	}

	if number == target {
		*flag = true
		output = "Congrats, you made the correct guess."
		return output
	} else if number < target {
		*flag = false
		output = "Your guess is low"
	} else if number > target {
		*flag = false
		output = "Your guess is high"
	}

	if math.Abs(float64(number-target)) <= 5 {
		output += " but close"
	}

	return output
}

func main() {
	var number int
	target := randomNumber()
	fmt.Println("Guess the target game starts")
	flag := false

	for guess := 0; guess < 5; guess++ {
		guessLeft(guess)
		fmt.Println("Enter number:")

		_, err := fmt.Scanf("%d\n", &number) //take guess number as an input from user.
		if err != nil {
			log.Fatal(err)
		}

		output := checkGuess(number, target, &flag)
		if flag {
			fmt.Println(output)
			break
		} else {
			fmt.Println(output)
		}
	}
	if !flag {
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
