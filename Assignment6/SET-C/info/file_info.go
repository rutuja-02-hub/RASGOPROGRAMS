package main

import (
	"fmt"
	"os"
)

func main() {
	// File name to check
	filename := "xml_handling/data.xml"

	// Get file info
	fileInfo, err := os.Stat(filename)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// Print file details
	fmt.Println("File Information:")
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
}
