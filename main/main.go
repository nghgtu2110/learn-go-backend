package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func prod(a, b int) int {
	return a * b
}

func main() {
	var a int
	fmt.Println("nhap vao so a: ")
	fmt.Scan(&a)
	var b int
	fmt.Println("nhap vao so b: ")
	fmt.Scan(&b)

	fmt.Printf("Tong cua 2 so a va b: %d\n", sum(a, b))
	fmt.Printf("Hieu cua 2 so a va b: %d\n", sub(a, b))
	fmt.Printf("Tich cua 2 so a va b: %d\n", prod(a, b))
}
