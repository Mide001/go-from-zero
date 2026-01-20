package main

import (
	"fmt"
	"time"
)

func mineBlock(id int, difficulty int) {
	fmt.Printf("Miner %d started mining...\n", id)
	time.Sleep(time.Duration(difficulty) * time.Millisecond * 100)
	fmt.Printf("Miner %d found a block!\n", id)
}

func example8_goroutines() {
	fmt.Println("\n=== Example 8: Goroutines (Concurrency) ===")

	for i := 1; i <= 3; i++ {
		go mineBlock(i, i*2)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("All miners finished")
}

func main() {
	example8_goroutines()
}
