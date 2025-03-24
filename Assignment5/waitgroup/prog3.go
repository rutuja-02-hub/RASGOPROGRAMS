package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500)
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	messages := []string{"Hello", "Goroutines", "are", "awesome!"}

	for _, msg := range messages {
		wg.Add(1)
		go printMessage(msg, &wg)
	}

	wg.Wait()
	fmt.Println("All messages printed.")
}
