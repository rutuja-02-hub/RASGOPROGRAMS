package main
import "fmt"

func main() {
    greet := func(name string) {
        fmt.Println("Hello,", name)
    }
    greet("Go Developer")
}
