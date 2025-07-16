package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("task start %d\n", id)
	time.Sleep(1 * time.Second)
	fmt.Println("task end", id)
}

func main() {
	start := time.Now()
	task(1)
	task(2)
	task(3)
	task(4)

	fmt.Println("Time to finish all tasks: ", time.Since(start))
}
