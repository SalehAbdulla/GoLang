package main

type Account struct {
	balance float64 // balance is private, we can't touch it from outside
}

// We must use Deposit or Balance, this is called encapsulation

// Exported method (capitalized)
func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

// Unexported field (lowercase) – hidden outside the package
func (a *Account) Balance() float64 {
	return a.balance
}

func main() {}
