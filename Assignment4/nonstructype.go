package main

import "fmt"

type Number int

func (n Number) isEven() bool {
    return n%2 == 0
}

func main() {
    var num Number = 10
    if num.isEven() {
        fmt.Println(num, "is even")
    } else {
        fmt.Println(num, "is odd")
    }
}
