package main

import (
	"fmt"
)

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(2, 10, 1, 4, 5, 90))
}
