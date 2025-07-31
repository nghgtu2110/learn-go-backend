package main

import (
	"fmt"
	"sync"
)

// use sync.Mutex Lock for clean code
func sum(sp *int, n int) {
	for i := 1; i <= n; i++ {
		*sp += 1
	}
}

func main() {

	s := 0

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("First sum ")
		sum(&s, 10000) // 100 000 will result in race condition
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		fmt.Println("Second sum ")
		sum(&s, 10000)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Total sum is: ", s)
}
