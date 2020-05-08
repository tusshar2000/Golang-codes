package main

import (
	"fmt"
)

func printInfo(s subs) {
	fmt.Println("Name", s.name)
	fmt.Println("Rate", s.rate)
	fmt.Println("Active?", s.active)
}

func main()
