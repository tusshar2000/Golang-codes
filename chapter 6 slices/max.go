package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max

}

func main() {
	fmt.Println(maximum(2, 3, 4))
	fmt.Println(maximum(2.1, 5.4, -19.2))
}
