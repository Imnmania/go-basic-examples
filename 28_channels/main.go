package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to channels...")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)

	// <-chan is RECEIVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// this is how you listen if the channel is open or not
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// chan<- is SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 5
		close(myCh)
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
