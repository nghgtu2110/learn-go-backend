package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	f := fibonacci
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Ten loi la:", err)
	} else {
		fmt.Println(f(n))
	}
}
