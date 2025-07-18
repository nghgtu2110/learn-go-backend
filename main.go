package main

import (
	"fmt"
	"others/charOccur"
	"others/personStruct"
	"others/readfilePkg"
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

	var array = [5]int{}
	array = [5]int{10, 5, -4, 2, 8}
	// array := []int{10, 5, -4, 2, 8}

	fmt.Println(slicePkg.Sum(array[:]))
	fmt.Println(slicePkg.Avg(array[:]))
	fmt.Println(slicePkg.FindMin(array[:]))
	fmt.Println(slicePkg.FindMax(array[:]))
	slicePkg.Sort(array[:])
	fmt.Println(array)

	var person personStruct.Person
	fmt.Scan(&person.Name, &person.Occupation, &person.BirthYear)

	fmt.Println(person.GetAge())
	fmt.Println(person.IsAppropriateCareer())

	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque vestibulum purus vitae laoreet bibendum. Proin mattis libero ac massa blandit."
	charMap := charOccur.CharacterOccurence(input)
	for k, v := range charMap {
		fmt.Printf("%c : %d,", k, v)
	}

	fmt.Println(readfilePkg.ReadPeopleFromFile("input.txt"))
}
