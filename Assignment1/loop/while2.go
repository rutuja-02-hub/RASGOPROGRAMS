package main
import "fmt"

func main() {
    var n, sum, i int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)

    i = 1
    for i <= n {
        sum += i
        i++
    }
    fmt.Println("Sum of first", n, "numbers is:", sum)
}
