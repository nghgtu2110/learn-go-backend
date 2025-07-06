package main

import (
	"fmt"
	"others/rectangle"
	"others/slicePkg"
	"others/stringPkg"
)

func main() {
	var lgth int
	fmt.Scan(&lgth)
	var wth int
	fmt.Scan(&wth)

	fmt.Println(rectangle.GetArea(lgth, wth))
	fmt.Println(rectangle.GetCircumference(lgth, wth))

	fmt.Println(stringPkg.IsEvenLengthString("hekfloanniag"))

	array := []int{10, 5, -4, 2, 8}

	fmt.Println(slicePkg.Sum(array))
	fmt.Println(slicePkg.Avg(array))
	fmt.Println(slicePkg.FindMin(array))
	fmt.Println(slicePkg.FindMax(array))
	slicePkg.Sort(array)
	fmt.Println(array)
}
