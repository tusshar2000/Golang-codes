package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 2) //buffered channel
	go readFromChannel(ch)
	go writeToChannel(ch)
	wg.Wait()
}

func writeToChannel(ch chan int) {
	fmt.Println("Starting tp write")
	time.Sleep(time.Second)
	ch <- 10
	ch <- 8
	ch <- 6
	// time.Sleep(time.Second) if this is run then beolow print statement wont be seen in o/p
	fmt.Println("data has been written to channel")
	wg.Done()
}

func readFromChannel(ch chan int) {
	fmt.Println("Starting to Read")
	value := <-ch
	value2 := <-ch
	fmt.Println("data in our channel is:", value, value2)
	wg.Done()
}
