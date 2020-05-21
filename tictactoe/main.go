package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	game()
}

func game() {

	board := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	startBoard(board)

	turn := 1
	gameOver := false
	player1, player2 := playerNameInput()

	for gameOver != true {
		if turn%2 == 1 {
			fmt.Printf("%s turn\n", player1)
		} else {
			fmt.Printf("%s turn\n", player2)
		}

		currentMove := askForPlay()
		board = executePlayerMove(currentMove, turn, board)
		result := resultAnalyzer(board)
		presentBoard(board)

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

//func to take input names of players.
func playerNameInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name of player 1:")
	player1, _ := reader.ReadString('\n')
	player1 = strings.TrimSpace(player1)

	fmt.Println("Enter name of player 2:")
	player2, _ := reader.ReadString('\n')
	player2 = strings.TrimSpace(player2)
	return player1, player2
}

//invoked only once at the starting to show user the type of array used (3x3).
func startBoard(board [9]int) {
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

//this func is used to print status of board after every correct move.
func presentBoard(board [9]int) {
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

//input reading func.
func askForPlay() int {
	fmt.Println("Select a move:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	moveInt, err := strconv.Atoi(text)
	//fmt.Println(err)
	//fmt.Println("1", err)
	for err != nil {
		fmt.Println("Please enter a number between 0-8.")
		fmt.Println("Select a move:")
		reader = bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
		moveInt, err = strconv.Atoi(text)
		//fmt.Println(err)
	}
	return moveInt
}

func executePlayerMove(moveInt int, turn int, b [9]int) [9]int {
	//out-of-bounds checker.
	//for loop performed until condition satisfied.
	for moveInt > 8 || moveInt < 0 {
		fmt.Println("Please enter a number between 0-8.")
		moveInt = askForPlay()
	}
	// Check for occupied spaces
	if b[moveInt] != 0 {
		fmt.Println("Please pick an unoccupied space.")
		moveInt = askForPlay()
		b = executePlayerMove(moveInt, turn, b)
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

func resultAnalyzer(b [9]int) int {
	sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	sums[0] = b[2] + b[4] + b[6]
	sums[1] = b[0] + b[3] + b[6]
	sums[2] = b[1] + b[4] + b[7]
	sums[3] = b[2] + b[5] + b[8]
	sums[4] = b[0] + b[4] + b[8]
	sums[5] = b[6] + b[7] + b[8]
	sums[6] = b[3] + b[4] + b[5]
	sums[7] = b[0] + b[1] + b[2]

	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 300 {
			return 2
		}
	}
	return 0
}
