package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "file_handling/sample.txt"

	// Open file in append mode, create if doesn't exist
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Content to append
	content := "\nThis is new content added at the end."

	// Write to file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content appended successfully.")
}
