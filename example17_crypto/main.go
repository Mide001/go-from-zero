package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Keypair struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func GenerateKeyPair() (*Keypair, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Keypair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
	}, nil
}

func (kp *Keypair) Sign(hash []byte) (r, s *big.Int, err error) {
	return ecdsa.Sign(rand.Reader, kp.PrivateKey, hash)
}

func (kp *Keypair) Verify(hash []byte, r, s *big.Int) bool {
	return ecdsa.Verify(kp.PublicKey, hash, r, s)
}

func example17_crypto() {
	fmt.Println("=== Example 17: ECDSA Signatures ===")

	keypair, err := GenerateKeyPair()
	if err != nil {
		fmt.Println("Error generating keypair: ", err)
		return
	}

	message := "Transfer 20 ETH to Bob"
	hash := sha256.Sum256([]byte(message))

	r, s, err := keypair.Sign(hash[:])
	if err != nil {
		fmt.Println("Error signing message: ", err)
		return
	}

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Hash: %s\n", hex.EncodeToString(hash[:]))
	fmt.Printf("Signature R: %s\n", r.String()[:20]+"...")
	fmt.Printf("Signature S: %s\n", s.String()[:20]+"...")

	valid := keypair.Verify(hash[:], r, s)
	fmt.Printf("Signature valid: %t\n", valid)

	wrongHash := sha256.Sum256([]byte("Wrong message"))
	valid = keypair.Verify(wrongHash[:], r, s)
	fmt.Printf("Wrong message valid: %t\n", valid)

	fmt.Println()
}

func main() {
	example17_crypto()
}
