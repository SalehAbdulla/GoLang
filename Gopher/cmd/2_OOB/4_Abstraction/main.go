package main

import "fmt"

// Abstraction -> Expose only what matters
type PaymentProcessor interface {
	Pay(amount float64) error
}

type PayPal struct{}

// Go’s interfaces achieve abstraction automatically.
func (p PayPal) Pay(amount float64) error {
	fmt.Println("Paid using PayPal")
	return nil
}
