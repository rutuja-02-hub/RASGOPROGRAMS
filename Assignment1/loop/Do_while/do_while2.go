package main
import "fmt"

func main() {
    var num, sum int
    sum = 0
    for {
        fmt.Print("Enter a number (negative to stop): ")
        fmt.Scan(&num)
        if num < 0 {
            break
        }
        sum += num
    }
    fmt.Println("Sum of entered numbers is:", sum)
}
