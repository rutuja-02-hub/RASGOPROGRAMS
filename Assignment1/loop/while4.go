package main
import "fmt"

func main() {
    var num, fact, i int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    fact = 1
    i = 1
    for i <= num {
        fact *= i
        i++
    }
    fmt.Println("Factorial of", num, "is:", fact)
}
