package main

import "fmt"

func validateTransaction(from string, to string, amount float64) error {
	if from == "" || to == "" {
		return fmt.Errorf("invalid address")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	if from == to {
		return fmt.Errorf("cannot send to yourself")
	}
	return nil
}

func example4_errors() {
	fmt.Println("\n=== Example 4: Error Handling ===")

	err := validateTransaction("alice", "bob", 5.0)
	if err != nil {
		fmt.Println("Validation failed: ", err)
	} else {
		fmt.Println("Validation passed")
	}

	err = validateTransaction("alice", "bob", 10.0)
	if err != nil {
		fmt.Println("Validation failed: ", err)
	}
}

func main() {
	example4_errors()
}
