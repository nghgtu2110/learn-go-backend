package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("task start %d\n", id)
	time.Sleep(1 * time.Second)
	fmt.Println("task end", id)
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
	fmt.Println("Time to finish all tasks: ", time.Since(start))
}
