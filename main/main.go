package main

import "fmt"

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int // Pointer variable
	intP = &i1    // Storing address of i1 in pointer variable
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

}
