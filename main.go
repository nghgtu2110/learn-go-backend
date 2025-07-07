package main

import "fmt"

func main() {
	s := []string{"Hello", "World", "from", "Go"}
	fmt.Println(s)
	x := s[:3]
	fmt.Println(x)
	y := s[2:]
	fmt.Println(y)

}
