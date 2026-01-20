package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transaction struct {
	From      string  `json:"from"`
	To        string  `json:"to"`
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
}

func example6_json() {
	fmt.Println("\n=== Example 6: JSON Handling ===")

	tx := Transaction{
		From:      "Alice",
		To:        "Bob",
		Amount:    10.5,
		Timestamp: time.Now().Unix(),
	}

	jsonData, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Println("JSON output:")
	fmt.Println(string(jsonData))

	jsonStr := `{"from":"Charlie", "to": "Dave", "amount": 5.0, "timestamp": 1234567890}`
	var newTx Transaction
	json.Unmarshal([]byte(jsonStr), &newTx)
	fmt.Printf("\nParsed: %s sent %.2f to %s\n", newTx.From, newTx.Amount, newTx.To)
}

func main() {
	example6_json()
}
