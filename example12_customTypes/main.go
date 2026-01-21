package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)

type Address string
type Hash [32]byte
type Wei *big.Int

func (a Address) IsValid() bool {
	return len(a) == 42 && strings.HasPrefix(string(a), "0x")
}

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

func NewWei(amount int64) Wei {
	return big.NewInt(amount)
}

func example12_customTypes() {
	fmt.Println("=== Example 12: Custom Types ===")

	addr := Address("0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb")
	fmt.Printf("Address: %s, Valid: %t\n", addr, addr.IsValid())

	data := "Genesis Block"
	hash := sha256.Sum256([]byte(data))
	var blockHash Hash = hash
	fmt.Printf("Block hash: %s\n", blockHash.String())

	oneEther := NewWei(1000000000000000000)
	fmt.Printf("1 ETH in wei: %s\n", (*big.Int)(oneEther).String())

	fmt.Println()
}

func main() {
	example12_customTypes()
}
