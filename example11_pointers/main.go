package main

import "fmt"

type Block struct {
	Data      [1000]byte
	Hash      string
	Timestamp int64
}

func example11_pointers() {
	fmt.Println("=== Example 11: Pointers & Memory ===")

	balance1 := 100.0
	balance2 := balance1
	balance2 = 200.0
	fmt.Printf("Value copy: balance1=%.2f, balance2=%.2f\n", balance1, balance2)

	walletBalance := 100.0
	ptr := &walletBalance
	*ptr = 200.0
	fmt.Printf("Pointer: walletBalance=%.2f, *ptr=%.2f\n", walletBalance, *ptr)

	modifyBlockValue := func(b Block) {
		b.Hash = "modified"
		fmt.Println("  Inside func (value):", b.Hash)
	}

	modifyBlockPointer := func(b *Block) {
		b.Hash = "modified"
		fmt.Println("  Inside func (pointer): ", b.Hash)
	}

	block := Block{Hash: "original"}
	fmt.Println("\nBefore:", block.Hash)
	modifyBlockValue(block)
	fmt.Println("After value call: ", block.Hash)

	modifyBlockPointer(&block)
	fmt.Println("After pointer call: ", block.Hash)

	fmt.Println()
}

func main() {
	example11_pointers()
}
