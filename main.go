package main

import (
	"fmt"
	"sync"
	"time"
)

// use sync.Mutex Lock for clean code
func sum(sp *int, n int) {
	for i := 1; i <= n; i++ {
		*sp += 1
	}
}

func main() {
	time_start := time.Now()
	s := 0

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time_start := time.Now()
		fmt.Println("First sum ")
		sum(&s, 100000) // 100 000 will result in race condition
		elapsed := time.Since(time_start)
		fmt.Println("Time taken: ", elapsed)
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		time_start := time.Now()
		fmt.Println("Second sum ")
		sum(&s, 100000)
		elapsed := time.Since(time_start)
		fmt.Println("Time taken: ", elapsed)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Total sum is: ", s)

	elapsed := time.Since(time_start)
	fmt.Println("Time taken: ", elapsed)
}
