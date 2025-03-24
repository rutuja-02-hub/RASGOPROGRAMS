package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the file name
	filename := "empty_file.txt"

	// Create an empty file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after creation

	fmt.Println("Empty file created successfully:", filename)
}
