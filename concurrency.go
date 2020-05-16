package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(runtime.GOMAXPROCS(10000))
	wg.Add(2)
	go delayedIteration1()
	go delayedIteration2()
	wg.Wait()
	// fmt.Scanln()
}

func delayedIteration1() {
	for i := 1; i < 6; i++ {
		time.Sleep(time.Second)
		fmt.Println("in 1:", i)
	}
	wg.Done()
}
func delayedIteration2() {
	for i := 1; i < 6; i++ {
		time.Sleep(time.Second)
		fmt.Println("in 2:", i)
	}
	wg.Done()
}
