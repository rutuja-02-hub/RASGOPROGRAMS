package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num % 3 == 0 && num % 5 == 0 {
        fmt.Println("The number is divisible by both 3 and 5.")
    } else {
        fmt.Println("The number is not divisible by both 3 and 5.")
    }
}
