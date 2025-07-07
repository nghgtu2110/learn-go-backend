package main

import "fmt"

func main() {
	var n uint64
	fmt.Scan(&n)

	fmt.Println(Factorial(n))
}

func Factorial(n uint64) (fac uint64) {
	fac = 1
	if n <= 1 {
		return
	} else {
		fac = (n * Factorial(n-1))
		return
	}
}
