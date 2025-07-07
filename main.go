package main

import "fmt"

func main() {

	arr := make([]int, 0, 10)

	fmt.Println("arr: ", arr, " len: ", len(arr), " cap: ", cap(arr))
	arr = append(arr, 10)
	arr = append(arr, 5)
	arr = append(arr, 1)
	arr = append(arr, 7)

	fmt.Println("arr: ", arr)
}
