package main

import (
	"fmt"
	"log"
	"os"

	action "reloaded/operation"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: expected go run . <inputFile> <outputFile>")
		// reading arguments from the command line
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	array := string(data)
	// converts the content from bytes to string.

	result := action.Operation(array) + "\n"
	outputData := []byte(result)
	err = os.WriteFile(outputFile, outputData, 0o644)
	if err != nil {
		panic(err)
	}
	// outFile := outputFile
}
