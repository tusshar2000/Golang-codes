package main

import (
	"fmt"
)

func main() {
	var number = 50
	f1()
	defer f2(number)
	number = 100
	fmt.Println("End of main. Number is:", number)
}

func f1() {
	defer fmt.Println("Hello from f1")
	fmt.Println("End of F1")
}

func f2(number int) {
	fmt.Println("End of F2. Number is:", number)
}
