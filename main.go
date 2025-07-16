package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		fmt.Println("Sending...")
	}()

	// Goroutine Anonymous
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println("Receiving...")
	}()

	// do anything on the channel isn't quick
	// so fmt.Println("") may complete sooner

	fmt.Println("Ending")
	time.Sleep(1 * time.Second)
}
