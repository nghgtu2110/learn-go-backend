package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
			fmt.Println("The input was:", inputString)
			fmt.Println("Read error:", readerError)
			break // EOF reached after printing the last line
		}
		//if readerError != nil {
		//	fmt.Println("Read error:", readerError)
		//	break
		//}
		fmt.Printf("The input was: %s", inputString)
	}
}
