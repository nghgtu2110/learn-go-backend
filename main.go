package main

import "fmt"

func main() {
	// function calls for multiple arg.
	fmt.Println(sumInts())
	fmt.Println(sumInts(2, 3))
	fmt.Println(sumInts(5, -2, 0, 9))
	fmt.Println(sumInts(0, 1, 2, 5, 10, 1, 2, -4))
}

// variadic function
func sumInts(list ...int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}
