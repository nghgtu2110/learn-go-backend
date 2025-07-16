package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting")
	ch := make(chan int)

	go func() {
		defer close(ch)
		ch <- 10
		ch <- 20
		ch <- 30
		fmt.Println("Sending...")
	}()

	// only work when there is close the channel statement
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Ending")
	time.Sleep(1 * time.Second)
}
