package main

import "fmt"

type Counter struct {
    value int
}

func (c *Counter) increment() {
    c.value++
}

func main() {
    c := Counter{value: 5}
    c.increment()
    fmt.Println("Counter value:", c.value)
}
