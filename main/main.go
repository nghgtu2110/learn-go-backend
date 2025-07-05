package main

import "fmt"

func main() {
	// Rectangular

	var a int
	var b int

	fmt.Print("Nhập a: ")
	fmt.Scan(&a)

	fmt.Print("Nhập b: ")
	fmt.Scan(&b)

	if a%b == 0 && b != 0 {
		fmt.Println("So a chia het cho so b")
	} else {
		fmt.Println("So a khong chia het cho so b")
	}

}
