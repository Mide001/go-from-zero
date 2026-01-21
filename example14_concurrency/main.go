package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

type Address string
type Hash [32]byte

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

type TxPool struct {
	transactions []Transaction
	mu           sync.Mutex
	newTxChan    chan Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make([]Transaction, 0),
		newTxChan:    make(chan Transaction, 100),
	}
}

func (pool *TxPool) AddTransaction(tx Transaction) {
	pool.newTxChan <- tx
}

func (pool *TxPool) Start() {
	go func() {
		for tx := range pool.newTxChan {
			pool.mu.Lock()
			pool.transactions = append(pool.transactions, tx)
			fmt.Printf("Added tx: %s -> %s (Pool size: %d)\n", tx.From, tx.To, len(pool.transactions))
			pool.mu.Unlock()
		}
	}()
}

func (pool *TxPool) GetPendingCount() int {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	return len(pool.transactions)
}

func example14_concurrency() {
	fmt.Println("=== Example 14: Channels & Goroutines ===")

	pool := NewTxPool()
	pool.Start()

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			tx := Transaction{
				From:  Address(fmt.Sprintf("0xUser%d", id)),
				To:    Address("0xRecipient"),
				Value: big.NewInt(int64(id * 100)),
				Nonce: uint64(id),
			}
			pool.AddTransaction(tx)
		}(i)
	}

	wg.Wait()
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("\nFinal pool size: %d\n", pool.GetPendingCount())
	fmt.Println()
}

func main() {
	example14_concurrency()
}
