package main

import "fmt"

func main() {
	s := []string{"Hello", "World", "from", "Go"}
	fmt.Println(s)
	x := s[:3]
	fmt.Println(x)
	y := s[2:]
	fmt.Println(y)

	fmt.Println("capacity of s:", cap(s))

	var a [6]int
	fmt.Println("capacity of a:", cap(a))

	arr := make([]int, 5, 10)

	fmt.Println("arr: ", arr, " len: ", len(arr), " cap: ", cap(arr))

}
