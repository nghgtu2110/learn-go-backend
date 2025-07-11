package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name       string
	Occupation string
	BirthYear  int
}

func (p Person) GetAge() int {
	return time.Now().Year() - p.BirthYear
}

func (p Person) IsAppropriateCareer() bool {
	return len(p.Name) >= 0 && (p.BirthYear%len(p.Name) == 0)
}

func main() {
	var person Person
	fmt.Scan(&person.Name, &person.Occupation, &person.BirthYear)

	fmt.Println(person.GetAge())
	fmt.Println(person.IsAppropriateCareer())
}
