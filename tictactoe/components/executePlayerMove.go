package components

import "fmt"

func ExecutePlayerMove(moveInt int, turn int, b [9]int) [9]int {
	//out-of-bounds checker.
	//for loop performed until condition satisfied.
	for moveInt > 8 || moveInt < 0 {
		fmt.Println("Please enter a number between 0-8.")
		moveInt = AskForPlay()
	}
	// Check for occupied spaces
	if b[moveInt] != 0 {
		fmt.Println("Please pick an unoccupied space.")
		moveInt = AskForPlay()
		b = ExecutePlayerMove(moveInt, turn, b)
	} else {
		/*player1 value is set to 1, for player2 value is set to 100.
		This difference in values is kept to avoid false positives.*/
		if turn%2 == 1 {
			b[moveInt] = 1
		} else {
			b[moveInt] = 100
		}
	}
	return b
}
