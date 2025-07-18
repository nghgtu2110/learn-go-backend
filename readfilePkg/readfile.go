package readfilePkg

import (
	"bufio"
	//"fmt"
	"io"
	"os"
	//"others/personStruct"
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
		line, lineError := inputReader.ReadString('\n')
		if lineError == io.EOF {
			parts := strings.Split(line, "|")
			//person := personStruct.Person{}
			for _, part := range parts {
				//fmt.Printf("Part %d: %s\n", i, part)
				output += part + "-"
			}
			break
		}

		line = strings.TrimSuffix(line, "\r\n")
		parts := strings.Split(line, "|")
		//person := personStruct.Person{}
		for _, part := range parts {
			//fmt.Printf("Part %d: %s\n", i, part)
			output += part + "-"
		}
	}
	return output
}
