package main
import "fmt"

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    if num % 5 == 0 {
        fmt.Println("The number is divisible by 5.")
    }
}
