package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type HashAlgorithm interface {
	Hash(data string) string
}

type SHA256Hasher struct{}

func (s SHA256Hasher) Hash(data string) string {
	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}

func example5_interfaces() {
	fmt.Println("\n=== Example 5: Interfaces ===")

	var hasher HashAlgorithm = SHA256Hasher{}

	data := "hello techwithmide"
	hash := hasher.Hash(data)
	fmt.Printf("Data: %s\nHash: %s\n", data, hash)
}

func main() {
	example5_interfaces()
}
