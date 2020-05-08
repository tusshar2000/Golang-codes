package main

import (
	"fmt"
)

type number int

func (n *number) Double() {
	*n *= 2
}

func main() {
	val := number(2)
	val.Double()
	fmt.Println(val)
}
