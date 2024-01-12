package main

import (
	"fmt"
	"os"
	"reboot"
)

func main() {
	// Checking the number of command-line arguments
	if len(os.Args) < 3 {
		fmt.Println("Input and output files are needed")
		return
	} else if len(os.Args) > 3 {
		fmt.Println("Only input and output files are needed")
		return
	}

	// Retrieving the input and output file paths from command-line arguments
	file := os.Args[1]
	resultFile := os.Args[2]

	// Reading the content of the input file
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Input file not found")
	}

	// Performing a conversion on the input text
	inputText := reboot.AllConversions(string(content))

	// Writing the converted text to the output file
	err = os.WriteFile(resultFile, []byte(inputText), 0666)
	if err != nil {
		fmt.Println("Error while writing to Output file")
	}
}
