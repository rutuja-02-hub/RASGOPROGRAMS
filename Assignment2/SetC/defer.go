package main
import (
    "fmt"
    "os"
)

func main() {
    
    file, err := os.Create("hello.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

   
    defer file.Close()

    _, err = file.WriteString("Hello, World!")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Println("File created and written successfully")
}
