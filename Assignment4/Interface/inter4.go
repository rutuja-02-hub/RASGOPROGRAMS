package main

import "fmt"

func execute(name string, callback func(string)) {
    callback(name)
}

func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    execute("Alice", greet)
}
