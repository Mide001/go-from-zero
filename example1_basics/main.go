package main

import (
	"fmt"
)

func example1_basics() {
	fmt.Println("=== Exmaple 1: Go Basics ===")

	var name string = "techwithmide"
	age := 40
	balance := 2100000000.0
	isActive := true

	fmt.Printf("Name: %s, Age: %d, Balance: %.2f, Active: %t\n", name, age, balance, isActive)

	const BLOCK_REWARD = 6.25
	fmt.Printf("Block Reward: %.2f BTC\n", BLOCK_REWARD)
}

func main() {
	example1_basics()
}
