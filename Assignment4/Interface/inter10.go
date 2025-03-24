package main

import "fmt"

func executeWithLog(action string, fn func()) {
	fmt.Println("Executing:", action)
	fn()
	fmt.Println("Finished:", action)
}

func main() {
	executeWithLog("Data Processing", func() {
		fmt.Println("Processing data...")
	})
}
