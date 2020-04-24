package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Game guess number")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guess := 0; guess <= 10; guess++ {
		fmt.Println("You have", 10-guess, "guesses left!")
		fmt.Println("Enter number")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if number < target {
			fmt.Println("Your guess is low")
		} else if number > target {
			fmt.Println("Your guess is high")
		} else {
			success = true
			fmt.Println("COngrats on winning the game.")
			break
		}
	}
	if !success {
		fmt.Println("sorry you couldnt guess, the number was:", target)
	}
}
