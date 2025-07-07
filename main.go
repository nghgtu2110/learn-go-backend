package main

import "fmt"

/*type flt func(int) bool // aliasing type*/

func isEven(n int) bool { // check if n is even or not
	return (n%2 == 0)
}

func filter(sl []int /*, f flt*/) (even, odd []int) { // split s into two slices: even and odd
	for _, v := range sl {
		if isEven(v) {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	even, odd := filter(slice)
	fmt.Println(even)
	fmt.Println(odd)
}
