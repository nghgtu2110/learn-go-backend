package main

import "fmt"

func abc(a int) {
	fmt.Println(a)
}

func main() {
	a := 10
	defer abc(a) // -> 10

	a += 10 // -> 20
	fmt.Println("hello", a) // a = 10
	
}
