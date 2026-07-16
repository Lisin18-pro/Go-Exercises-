package main

import (
	"fmt"
	"time"
)

func main() {
	var seconds int

	fmt.Print("Enter countdown time (in seconds): ")
	fmt.Scan(&seconds)

	fmt.Println("\nStarting countdown...")

	for i := seconds; i > 0; i-- {
		fmt.Printf("%d seconds remaining...\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("⏰ Time's Up!")
}
