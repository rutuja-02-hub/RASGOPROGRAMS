package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fetching data...")
	time.Sleep(2 * time.Second)
	fmt.Println("Data fetched.")
}

func processData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing data...")
	time.Sleep(1 * time.Second)
	fmt.Println("Data processed.")
}

func saveData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Saving data...")
	time.Sleep(1 * time.Second)
	fmt.Println("Data saved.")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go fetchData(&wg)
	go processData(&wg)
	go saveData(&wg)

	wg.Wait()
	fmt.Println("All tasks completed.")
}
