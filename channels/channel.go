//unbuffered channel
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go writeToChannel(ch)
	value := <-ch
	fmt.Println("data in our channel is:", value)
}

func writeToChannel(ch chan int) {
	fmt.Println("Inside Go routine")
	ch <- 10
	// time.Sleep(time.Second) if this is run then beolow print statement wont be seen in o/p
	fmt.Println("data has been written to channel")
}
