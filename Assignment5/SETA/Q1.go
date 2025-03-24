package main

import (
	"fmt"
	"strconv"
)

func calculateSquaresAndCubes(num int, squaresChan chan int, cubesChan chan int) {
	squaresSum, cubesSum := 0, 0
	numStr := strconv.Itoa(num)

	for _, digit := range numStr {
		n, _ := strconv.Atoi(string(digit))
		squaresSum += n * n
		cubesSum += n * n * n
	}

	squaresChan <- squaresSum
	cubesChan <- cubesSum
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	squaresChan := make(chan int)
	cubesChan := make(chan int)

	go calculateSquaresAndCubes(num, squaresChan, cubesChan)

	squaresSum := <-squaresChan
	cubesSum := <-cubesChan

	fmt.Printf("Sum of squares = %d\n", squaresSum)
	fmt.Printf("Sum of cubes = %d\n", cubesSum)
	fmt.Printf("Final sum of squares and cubes = %d\n", squaresSum+cubesSum)
}
