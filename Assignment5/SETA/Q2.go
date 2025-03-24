package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNumbers(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(time.Duration(rand.Intn(251)) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go generateNumbers(i, &wg)
	}

	wg.Wait()
}
