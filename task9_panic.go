package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Close resource")
		details := recover()
		fmt.Println("Details", details)
	}()
	fmt.Println("Open Resouce")
	panic("panicking: I don't know what to do")
	recover()

}
