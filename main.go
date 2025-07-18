package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	inputFile, inputError := os.Open("input.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerError := inputReader.ReadString('|')
		if readerError == io.EOF {
			fmt.Printf("The input was: %s\n", inputString)
			break // EOF reached after printing the last line
		}
		// Remove the trailing newline '\n' or any delimiter
		inputString = strings.TrimSuffix(inputString, "|")
		// Use TrimRight if you want to remove \r\n (Windows)
		// line = strings.TrimRight(line, "\r\n")
		fmt.Printf("The input was: %s\n", inputString)
	}
}
