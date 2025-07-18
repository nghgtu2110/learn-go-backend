package readfilePkg

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func PrintFromFile(filepath string) string {
	inputFile, inputError := os.Open(filepath)
	if inputError != nil {
		return "An error occurred on opening the input file\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n"
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	output := ""

	for {
		inputString, readerError := inputReader.ReadString('|')
		if readerError == io.EOF {
			output += "The input was:" + inputString + "\n"
			break
		}

		inputString = strings.TrimSuffix(inputString, "|")

		output += "The input was:" + inputString + "\n"
	}
	return output
}
