package main
import "fmt"

func main() {
    var n, i int
    fmt.Print("Enter a number: ")
    fmt.Scan(&n)

    i = 1
    for i <= n {
        if i % 2 == 0 {
            fmt.Println(i)
        }
        i++
    }
}
