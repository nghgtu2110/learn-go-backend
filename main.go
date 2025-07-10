package main

import "fmt"

func main() {
	var value int
	var isPresent bool
	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Paris"] = 20
	map1["Washington"] = 25

	// checking existence of a key
	if value, isPresent = map1["Beijing"]; isPresent {
		fmt.Printf("The value of Beijing in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Beijing")
	}

	// checking existence of a key
	if value, isPresent = map1["Paris"]; isPresent {
		fmt.Printf("The value of Paris in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Paris")
	}

	// delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"] // checking existence of a key

	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}
}
