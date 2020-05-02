package main

import "fmt"

func outer() func(int, int) float64 {
	fmt.Println("start")
	var innerFunction = func(number1, number2 int) float64 {
		return float64(number1) / float64(number2)
	}
	return innerFunction
}

func main() {
	test := outer()
	fmt.Println("break")
	fmt.Println(test(2, 4))
}
