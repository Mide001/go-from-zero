package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

type Hash [32]byte

type Datebase interface {
	Put(key []byte, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) error
	Has(key []byte) bool
}

type MemoryDB struct {
	data map[string][]byte
	mu   sync.RWMutex
}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{
		data: make(map[string][]byte),
	}
}

func (db *MemoryDB) Put(key []byte, value []byte) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.data[string(key)] = value
	return nil
}

func (db *MemoryDB) Get(key []byte) ([]byte, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	if val, exists := db.data[string(key)]; exists {
		return val, nil
	}
	return nil, fmt.Errorf("Key not found")
}

func (db *MemoryDB) Delete(key []byte) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	delete(db.data, string(key))
	return nil
}

func (db *MemoryDB) Has(key []byte) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()
	_, exists := db.data[string(key)]
	return exists
}

func storeBlock(db Datebase, blockNum uint64, hash Hash) error {
	key := []byte(fmt.Sprintf("block:%d", blockNum))
	value := hash[:]
	return db.Put(key, value)
}

func example15_interfaces() {
	fmt.Println("=== Example 15: Interface Abstraction ===")

	var db Datebase = NewMemoryDB()

	hash := sha256.Sum256([]byte("Block 1 data"))
	var blockHash Hash = hash

	storeBlock(db, 1, blockHash)

	val, err := db.Get([]byte("block:1"))
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Stored block hash: %s\n", hex.EncodeToString(val))
	}

	fmt.Printf("Has block 1: %t\n", db.Has([]byte("block:1")))
	fmt.Printf("Has block 2: %t\n", db.Has([]byte("block:2")))

	fmt.Println()
}

func main() {
	example15_interfaces()
}
