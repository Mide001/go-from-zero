package main

import (
	"fmt"
	"math/big"
)

var (
	ErrInsufficientBalance = fmt.Errorf("insufficient balance")
	ErrInvalidNonce        = fmt.Errorf("invalid nonce")
	ErrGasTooLow           = fmt.Errorf("gas limit too low")
)

type Address string
type Hash [32]byte

type ValidationError struct {
	Field   string
	Message string
}

type Transaction struct {
	From      Address
	To        Address
	Value     *big.Int
	Nonce     uint64
	GasLimit  uint64
	GasPrice  *big.Int
	Data      []byte
	Timestamp int64
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}

func validateTx(tx Transaction, balance *big.Int, expectedNonce uint64) error {
	if tx.Value.Cmp(balance) > 0 {
		return ErrInsufficientBalance
	}

	if tx.Nonce != expectedNonce {
		return fmt.Errorf("%w: expected %d, got %d", ErrInvalidNonce, expectedNonce, tx.Nonce)
	}

	if tx.GasLimit < 21000 {
		return &ValidationError{
			Field:   "GasLimit",
			Message: "must be at least 21000",
		}
	}

	return nil
}

func example16_errors() {
	fmt.Println("=== Example 16: Error Handling ===")

	tx := Transaction{
		From:     "0xAlice",
		To:       "0xBob",
		Value:    big.NewInt(1000),
		Nonce:    2,
		GasLimit: 21000,
	}

	balance := big.NewInt(500)
	expectedNonce := uint64(1)

	err := validateTx(tx, balance, expectedNonce)
	if err != nil {
		if err == ErrInsufficientBalance {
			fmt.Println("Transaction failed: Not enough funds")
		} else if valErr, ok := err.(*ValidationError); ok {
			fmt.Printf("Validation failed: %s\n", valErr.Error())
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	} else {
		fmt.Println("Transaction valid")
	}

	fmt.Println()
}

func main() {
	example16_errors()
}
