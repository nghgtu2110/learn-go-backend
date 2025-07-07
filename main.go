package main

import "fmt"

func SumProductDiff(i, j int) (s int, p int, d int) {
	s = i + j
	p = i * j
	d = i - j
	return s, p, d
}
func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(SumProductDiff(a, b))
}
