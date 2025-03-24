package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			fmt.Printf("Worker %d is processing...\n", workerID)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d is done\n", workerID)
		}(i)
	}

	wg.Wait()
	fmt.Println("All workers finished execution.")
}
