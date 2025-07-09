package main

import (
	"fmt"
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

	for {
		fmt.Print("Enter the length of rectangle: ")
		_, l_err := fmt.Scanf("%f", &rectangle.Length)
		if l_err != nil {
			fmt.Println(l_err)
			fmt.Println("The value entered is invalid. Please try again.")
		}
		if l_err == nil && rectangle.Length > 0 {
			break
		}
	}

	for {
		fmt.Print("Enter the width of rectangle: ")
		_, l_err := fmt.Scanf("%f", &rectangle.Width)
		if l_err != nil {
			fmt.Println(l_err)
			fmt.Println("The value entered is invalid. Please try again.")
		}
		if l_err == nil && rectangle.Width > 0 {
			break
		}
	}

	fmt.Println(rectangle.Perimeter())
	fmt.Println(rectangle.Area())
}
