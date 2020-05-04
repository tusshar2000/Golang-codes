package main

import "fmt"

func addition(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func subtraction(number1, number2 float64) float64 {
	return number1 - number2
}

func multiplication(number1, number2 float64) float64 {
	return number1 * number2
}

func division(number1, number2 float64) float64 {
	return number1 / number2
}

func mathOperation(functionName func(float64, float64) float64) float64 {
	return functionName(2.2, 3.3)

}

//another way -->

// func mathOperation(functionName func(float64, float64) float64, number1 float64, number2 float64) float64 {
// 	return functionName(2.2, 3.3)
// }

func main() {
	fmt.Println(mathOperation(subtraction))
	// fmt.Println(subtraction(2.2, 3.3))
	// fmt.Println(multiplication(2.2, 3.3))
	// fmt.Println(division(2.2, 3.3))
}
