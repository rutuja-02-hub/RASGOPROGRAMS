package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num >= 1 && num <= 100 {
        fmt.Println("The number is within the range of 1 to 100.")
    } else {
        fmt.Println("The number is outside the range of 1 to 100.")
    }
}
