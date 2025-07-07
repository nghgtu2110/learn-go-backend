package main

import (
	"fmt"
)

func func1(s string) (n int) {

	defer func() {
		fmt.Println("s :", s)
	}()

	return 7
}

func main() {
	fmt.Println(func1("Go"))
}
