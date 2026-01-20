package main

import "fmt"

type Wallet struct {
	Address string
	Balance float64
	Nonce   int
}

func (w *Wallet) Deposit(amount float64) {
	w.Balance += amount
	fmt.Printf("Deposited %.2f to %s. New Balance: %.2f\n", amount, w.Address, w.Balance)
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount > w.Balance {
		return fmt.Errorf("insufficient funds: need %.2f but have %.2f", amount, w.Balance)
	}

	w.Balance -= amount
	w.Nonce++
	return nil
}

func example3_structs() {
	fmt.Println("\n=== Example 3: Structs & Methods ===")

	wallet := Wallet{
		Address: "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
		Balance: 10.0,
		Nonce:   0,
	}

	wallet.Deposit(5.0)

	err := wallet.Withdraw(3.0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Withdrawal successful. Balance: %.2f, Nonce: %d\n", wallet.Balance, wallet.Nonce)
	}
}

func main() {
	example3_structs()
}
