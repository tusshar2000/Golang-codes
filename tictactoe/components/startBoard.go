package components

import "fmt"

//invoked only once at the starting to show user the type of array used (3x3).
func StartBoard(board [9]int) {
	for i, v := range board {
		if v == 0 {
			fmt.Printf("%d", i)
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
