package main

import (
	"fmt"
)

func divideArrayElement(array [5]int) [5]float64 {
	var answer [5]float64
	for i := range array {
		answer[i] = float64(array[i]) / 3.0
	}
	return answer
}
func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(divideArrayElement(numbers))

}
