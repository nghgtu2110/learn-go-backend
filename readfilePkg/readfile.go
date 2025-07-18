package readfilePkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"others/personStruct"
	"strconv"
	"strings"
)

func ReadPeopleFromFile(filepath string) []personStruct.Person {
	var people []personStruct.Person
	inputFile, inputError := os.Open(filepath)
	if inputError != nil {
		fmt.Println("An error occurred on opening the input file\n",
			"Does the file exist?\n",
			"Have you got access to it?")
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		line, lineError := inputReader.ReadString('\n')
		if lineError == io.EOF {
			parts := strings.Split(line, "|")

			person := personStruct.Person{}
			person.Name = parts[0]
			person.Occupation = parts[1]

			birthYear, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("Error converting birth year:", err)
				break
			}

			person.BirthYear = birthYear
			people = append(people, person)
			break
		}

		line = strings.TrimSuffix(line, "\r\n")
		parts := strings.Split(line, "|")

		person := personStruct.Person{}
		person.Name = parts[0]
		person.Occupation = parts[1]

		birthYear, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Error converting birth year:", err)
			break
		}
		person.BirthYear = birthYear

		people = append(people, person)
	}
	return people
}
