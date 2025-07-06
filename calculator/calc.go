package calculator

func Add(a int, b int) int { // implicit named return output
	return a + b
}

func Multi(a, b int) (c int) { // explicit named return output
	c = a * b
	return
}

func AddAndCheckNegative(a, b int) (int, bool) {
	sum := a + b
	isNegative := sum < 0
	return sum, isNegative
}
