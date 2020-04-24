package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, val := range numbers {
		sum += val
	}
	denominator := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/denominator)
}
