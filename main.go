package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 7, 9, 6, 5, 10, 25}

	for index, elem := range arr {
		fmt.Println("At index: ", index, " element is: ", elem)
	}

	for index := range arr {
		fmt.Println("At index: ", index, " element is: ", arr[index])
	}

	for _, elem := range arr {
		fmt.Println("Element is: ", elem)
	}
}
