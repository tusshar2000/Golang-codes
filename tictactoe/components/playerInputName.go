package components

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func to take input names of players.
func PlayerNameInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter name of player 1:")
	player1, _ := reader.ReadString('\n')
	player1 = strings.TrimSpace(player1)

	fmt.Println("Enter name of player 2:")
	player2, _ := reader.ReadString('\n')
	player2 = strings.TrimSpace(player2)
	return player1, player2
}
