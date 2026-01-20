package main

import (
	"fmt"
)

func example2_collections() {
	fmt.Println("\n=== Example 2: Slices & Maps ===")

	transactions := []string{}
	transactions = append(transactions, "Alice->Bob: 5 ETH")
	transactions = append(transactions, "Bob->Charlie: 2 ETH")
	transactions = append(transactions, "Charlie->Alice: 1 ETH")

	fmt.Println("Transactions: ")
	for i, tx := range transactions {
		fmt.Printf("  %d: %s\n", i, tx)
	}

	balances := make(map[string]float64)
	balances["Alice"] = 100.0
	balances["Bob"] = 50.0
	balances["Charlie"] = 75.0

	fmt.Println("\nBalances: ")
	for account, balance := range balances {
		fmt.Printf(". %s: %.2f ETH\n", account, balance)
	}

	if bal, exists := balances["Dave"]; exists {
		fmt.Printf("Dave's balance: %.2f\n", bal)
	} else {
		fmt.Println("  Dave has no balance")
	}
}

func main() {
	example2_collections()
}
