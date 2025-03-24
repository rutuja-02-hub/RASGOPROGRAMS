package stringsutil

import "strings"

// ToUpper converts a string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// CountWords counts the number of words in a string
func CountWords(s string) int {
	return len(strings.Fields(s))
}
