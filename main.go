package main

import "fmt"

func DoSomething1(a *int) *int {
	b := a
	return b
}
func DoSomething2(a int) *int {
	b := &a
	return b
}

func main() {
	x := 10
	//x_ptr := &x
	fmt.Println(DoSomething1(&x))
	fmt.Println(DoSomething2(x))
}
