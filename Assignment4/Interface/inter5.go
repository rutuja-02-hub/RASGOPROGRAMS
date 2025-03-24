package main

import (
    "fmt"
    "sort"
)

func sortNumbers(numbers []int, comparator func(i, j int) bool) {
    sort.Slice(numbers, comparator)
}

func main() {
    nums := []int{4, 2, 9, 1, 5}
    
    sortNumbers(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })

    fmt.Println("Sorted Numbers:", nums)
}
