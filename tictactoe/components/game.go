package components

import "fmt"

func Game() {

	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	StartBoard(board)

	turn := 1
	gameOver := false
	player1, player2 := PlayerNameInput()

	for gameOver != true {
		if turn%2 == 1 {
			fmt.Printf("%s turn\n", player1)
		} else {
			fmt.Printf("%s turn\n", player2)
		}

		currentMove := AskForPlay()
		board = ExecutePlayerMove(currentMove, turn, board)
		result := ResultAnalyzer(board)
		PresentBoard(board)

		if result > 0 {
			fmt.Printf("Player %d wins!\n\n", result)
			gameOver = true
		} else if turn == 9 {
			fmt.Printf("Tie game!\n\n")
			gameOver = true
		} else {
			turn++
		}

	}
}
