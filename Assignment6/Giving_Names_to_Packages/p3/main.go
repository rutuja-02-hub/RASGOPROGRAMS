package main

import (
	"Assignment6/stringsutil"
	"fmt"
)

func main() {
	text := "hello world from golang"

	// Using stringsutil package
	uppercaseText := stringsutil.ToUpper(text)
	wordCount := stringsutil.CountWords(text)

	// Print results
	fmt.Println("Uppercase:", uppercaseText)
	fmt.Println("Word Count:", wordCount)
}
