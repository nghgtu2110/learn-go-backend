package calculator

func Add(a int, b int) int { // implicit named return output
	return a + b
}

func Multi(a, b int) (c int) { // explicit named return output
	c = a * b
	return
}
