package main

import "fmt"

func main() {
	s := "good bye"
	var p *string = &s
	*p = "ciao" // changing the value at &s

	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
}
