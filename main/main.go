package main

import "fmt"

func main() {
	var a int
	fmt.Println("nhap vao so a: ")
	fmt.Scan(&a)
	var b int
	fmt.Println("nhap vao so b: ")
	fmt.Scan(&b)

	fmt.Printf("Tong cua 2 so a va b: %d\n", a+b)
	fmt.Printf("Hieu cua 2 so a va b: %d\n", a-b)
	fmt.Printf("Tich cua 2 so a va b: %d\n", a*b)
}
