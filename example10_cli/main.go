package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func example10_cli() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		fmt.Println("Commands: hash <text>, balance <address>")
		return
	}

	command := os.Args[1]

	switch command {
	case "hash":
		if len(os.Args) < 3 {
			fmt.Println("Usage: hash <text>")
			return
		}
		text := os.Args[2]
		h := sha256.Sum256([]byte(text))
		fmt.Printf("Hash: %s\n", hex.EncodeToString(h[:]))

	case "balance":
		if len(os.Args) < 3 {
			fmt.Println("Usage: balance <address>")
			return
		}
		address := os.Args[2]

		balances := map[string]float64{
			"alice": 100.0,
			"bob":   50.0,
		}

		if bal, exists := balances[address]; exists {
			fmt.Printf("Balance for %s: %.2f ETH\n", address, bal)
		} else {
			fmt.Printf("No balance found for %s\n", address)
		}
	default:
		fmt.Println("Unknown command: ", command)
	}
}

func main() {
	example10_cli()
}
