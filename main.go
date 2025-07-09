package main

import (
	"fmt"
	"os"
)

type Rectangle struct {
	Length float64 `desc:"Length of rectangle coordinates"`
	Width  float64 `desc:"Width of rectangle coordinates"`
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func main() {

	var rectangle Rectangle

	_, l_err := fmt.Scanf("%f", &rectangle.Length)
	if l_err != nil {
		fmt.Println(l_err)
		os.Exit(1)
	}

	_, w_err := fmt.Scanf("%f", &rectangle.Width)
	if w_err != nil {
		fmt.Println(w_err)
		os.Exit(1)
	}

	fmt.Println(rectangle.Perimeter())
	fmt.Println(rectangle.Area())
}
