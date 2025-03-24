package main

import "fmt"

func main() {
    var str1, str2 string

    // Accept two strings from the user
    fmt.Print("Enter the first string: ")
    fmt.Scan(&str1)
    fmt.Print("Enter the second string: ")
    fmt.Scan(&str2)

    // Compare the two strings
    if str1 == str2 {
        fmt.Println("The strings are equal.")
    } else {
        fmt.Println("The strings are not equal.")
    }
}
