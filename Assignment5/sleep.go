package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program started")
	time.Sleep(2 * time.Second)
	fmt.Println("2 seconds later...")
	time.Sleep(3 * time.Second)
	fmt.Println("5 seconds later... Program finished")
}
