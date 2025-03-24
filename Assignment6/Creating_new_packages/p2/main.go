package main

import (
	"Assignment6/stringutils"
	"fmt"
)

func main() {
	str := "hello"

	// Using stringutils package
	upper := stringutils.ToUpperCase(str)
	reversed := stringutils.ReverseString(str)

	// Print results
	fmt.Println("Uppercase:", upper)
	fmt.Println("Reversed:", reversed)
}
