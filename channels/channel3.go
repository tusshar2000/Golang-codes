package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)    //buffered channel
	ch2 := make(chan struct{}) //signal channel, doesnt require any memory
	go readFromChannel(ch)
	go writeToChannel(ch)
	wg.Wait()
	close(ch) //we cannot read or write after a channel is closed.
	go func() {
		ch2 <- struct{}{} //0 value memory allocation
	}()
	for i := range ch {
		fmt.Println("Inside range", i) //for this to run we need channel to be closed
	}

	for i := 0; i < 2; i++ { //we have 2 channels
		select {
		case chan1, ok := <-ch:
			fmt.Println("channel 1 has data", chan1, ok) //returns false because we closed it above
		case chan2, ok := <-ch2:
			fmt.Println("channel 2 has data", chan2, ok) //returns truue because we didnt close the channel.

			// default can also be used oin select.
		}
	}
}

func writeToChannel(ch chan<- int) {
	fmt.Println("Starting tp write")
	time.Sleep(time.Second)
	ch <- 10
	ch <- 8
	ch <- 6
	//value:=ch this wont work as it send only channel
	// time.Sleep(time.Second) if this is run then beolow print statement wont be seen in o/p
	fmt.Println("data has been written to channel")
	wg.Done()
}

func readFromChannel(ch <-chan int) {
	fmt.Println("Starting to Read")
	value, ok := <-ch //ok indicates channel is opended or closed
	value2 := <-ch
	fmt.Println("data in our channel is:", value, value2, ok)
	wg.Done()
}
