package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("Matrix Elements:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("Element [%d][%d]: %d\n", i, j, matrix[i][j])
		}
	}
}
