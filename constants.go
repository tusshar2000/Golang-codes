package main

import "fmt"

const (
	cat   = iota
	dog   = iota
	camel = iota
	horse = iota
)
const (
	cat2 int = iota + 1
	dog2
	camel2
	horse2
)

func main() {
	const myConstant = 6 //cannot manipulate and should avoid shadowing.
	fmt.Println(cat, dog, camel, horse)
	fmt.Println(cat2, dog2, camel2, horse2)
}
