package main

import "fmt"

func function1() func() {
	fmt.Println("Hello")
	var function2 = func() {
		fmt.Println("Namaste")
	}
	return function2
}

func main() {
	outer := function1()
	fmt.Println("break")
	outer()
}
