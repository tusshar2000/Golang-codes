package components

import "fmt"

//this func is used to print status of board after every correct move.
func PresentBoard(board [9]int) {
	for i, v := range board {
		if v == 0 {
			fmt.Printf(" ")
		} else if v == 1 {
			fmt.Printf("X")
		} else if v == 100 {
			fmt.Printf("O")
		}
		// And now the decorations
		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}
