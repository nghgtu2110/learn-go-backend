package main

import "fmt"

func main() {
	// error of pointers and addresses:

	const i = 5
	ptr1 := &i  //error: cannot take the address of i
	ptr2 := &10 //error: cannot take the address of 10

	fmt.Printf("Here is the pointer p: %p\n", ptr1) // cannot print
	fmt.Printf("Here is the pointer p: %p\n", ptr2) // cannot print
}
