package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

type Address string
type Hash [32]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
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

type SignedTransaction struct {
	Transaction
	V, R, S *big.Int
}

func (st *SignedTransaction) Verify() bool {
	return st.V != nil && st.R != nil && st.S != nil
}

func (st *SignedTransaction) Hash() Hash {
	data := fmt.Sprintf("%s%s%s%d", st.From, st.To, st.Value, st.Nonce)
	return sha256.Sum256([]byte(data))
}

func example13composition() {
	fmt.Printf("=== Example 13: Struct Composition ===\n")

	tx := SignedTransaction{
		Transaction: Transaction{
			From:      "0xmide",
			To:        "0xbob",
			Value:     big.NewInt(1000000),
			Nonce:     1,
			GasLimit:  21000,
			GasPrice:  big.NewInt(20000000000),
			Timestamp: time.Now().Unix(),
		},
		V: big.NewInt(27),
		R: big.NewInt(12345),
		S: big.NewInt(67890),
	}

	fmt.Printf("From: %s\n", tx.From)
	fmt.Printf("To: %s\n", tx.To)
	fmt.Printf("Value: %s wei\n", tx.Value.String())
	fmt.Printf("Valid signature: %t\n", tx.Verify())
	fmt.Printf("Tx Hash: %s\n", tx.Hash().String())

	fmt.Println()
}

func main() {
	example13composition()
}
