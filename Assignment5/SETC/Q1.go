package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 5

func worker(id int, wg *sync.WaitGroup, barrier *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Assembling details...\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d: Finished assembling, waiting at checkpoint...\n", id)

	barrier.Done()
	barrier.Wait()

	fmt.Printf("Worker %d: Proceeding to the next task.\n", id)
}

func main() {
	var wg sync.WaitGroup
	var barrier sync.WaitGroup

	wg.Add(numWorkers)
	barrier.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, &barrier)
	}

	wg.Wait()
}
