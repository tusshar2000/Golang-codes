package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("can't get square root of a negative number")
	}
	return math.Sqrt(num), nil
}

func main() {
	root, err := sqrt(-9.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%0.3f", root)
	}
}
