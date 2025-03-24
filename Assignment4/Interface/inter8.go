package main

import "fmt"

func filter(nums []int, condition func(int) bool) []int {
    var result []int
    for _, num := range nums {
        if condition(num) {
            result = append(result, num)
        }
    }
    return result
}

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}
    evens := filter(numbers, func(n int) bool { return n%2 == 0 })
    fmt.Println("Even Numbers:", evens)
}
