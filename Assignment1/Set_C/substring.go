package main

import "fmt"

// Function to check if the first string is a substring of the second string
func isSubstring(str1, str2 string) bool {
    // Use the built-in "Contains" function from the "strings" package to check if str1 is a substring of str2
    return contains(str2, str1)
}

// Function to check if a string contains another string
func contains(str1, str2 string) bool {
    return len(str2) > 0 && len(str1) > 0 && strings.Contains(str1, str2)
}

func main() {
    var str1, str2 string

    // Accept two strings from the user
    fmt.Print("Enter the first string: ")
    fmt.Scan(&str1)
    fmt.Print("Enter the second string: ")
    fmt.Scan(&str2)

    // Check if the first string is a substring of the second string
    if contains(str2, str1) {
        fmt.Println("The first string is a substring of the second string.")
    } else {
        fmt.Println("The first string is NOT a substring of the second string.")
    }
}
