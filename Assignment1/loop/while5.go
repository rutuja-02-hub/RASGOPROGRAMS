package main
import "fmt"

func main() {
    var num, i int
    var isPrime bool = true
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)

    i = 2
    for i < num {
        if num % i == 0 {
            isPrime = false
            break
        }
        i++
    }

    if isPrime && num > 1 {
        fmt.Println(num, "is a prime number.")
    } else {
        fmt.Println(num, "is not a prime number.")
    }
}
