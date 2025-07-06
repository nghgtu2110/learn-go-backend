package main

import "fmt"

func main() {
	max := 4
	if val := 10; val > max {
		fmt.Println("Bigger than max")
	}

	fmt.Println("Val is undefined here: " + val)
}
