package main

import (
	"fmt"
)

func variadicNumbers(numbers ...int) {
	fmt.Println(numbers)
	println(numbers)

}

func main() {
	variadicNumbers(1, 2, 3, 4, 5)

}
