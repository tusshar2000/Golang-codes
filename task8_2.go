package main

import (
	// "errors"
	"fmt"
)

type MyError struct {
	s string
}

func (err *MyError) Error() string {
	return err.s
}
func main() {
	quotient, err := divide(10, 0)
	if err != nil {
		fmt.Println("error occured", err)
		return
	}
	fmt.Println("result is", quotient)
}

func divide(numerator, denomintaor int) (int, error) {
	if denomintaor == 0 {
		return 0, &MyError{s: "cannot divide by zero"} //errors.New("cannot divide by zero")
	}
	return numerator / denomintaor, nil
}
