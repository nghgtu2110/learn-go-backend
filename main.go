package main

import "fmt"

func main() {
	s := []struct {
		a int
		b bool
	}{
		{a: 1, b: true},
		{2, false},
	}

	fmt.Println(s)
}
