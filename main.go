package main

import (
	"fmt"
	"math"
)

func compute(fn func(x, y float64) float64, x, y float64) float64 {
	return fn(x, y)
}

func mathpow(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(compute(hypot, 5, 12))
	fmt.Println(mathpow(math.Pow))
}
