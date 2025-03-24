package main

import "fmt"

func main() {
    matrix := [2][3]int{
        {1, 2, 3},
        {4, 5, 6},
    }

    fmt.Println("Matrix Elements:")
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            fmt.Printf("Element [%d][%d]: %d\n", i, j, matrix[i][j])
        }
    }
}
