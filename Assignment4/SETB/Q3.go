package main

import "fmt"

type ArrayCopy struct {
    source      [5]int
    destination [5]int
}

func (a *ArrayCopy) copyArray() {
    for i := 0; i < len(a.source); i++ {
        a.destination[i] = a.source[i]
    }
}

func (a ArrayCopy) showArrays() {
    fmt.Println("Source Array:", a.source)
    fmt.Println("Destination Array:", a.destination)
}

func main() {
    arr := ArrayCopy{source: [5]int{1, 2, 3, 4, 5}}
    arr.copyArray()
    arr.showArrays()
}
