package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"hoten"` /*if there is space between json: and "...", it won't work */
	Age  int    `json:"tuoi"`
}

func main() {
	nobody := Person{
		Name: "Nobody",
		Age:  40,
	}

	json, err := json.Marshal(nobody)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(nobody)
	fmt.Println(string(json))
}
