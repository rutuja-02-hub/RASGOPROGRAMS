package main

import "fmt"

func main() {
    i := 5
    for {
        fmt.Println(i)
        i--
        if i < 1 {
            break
        }
    }
}
