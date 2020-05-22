package components

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//input reading func.
func AskForPlay() int {
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
