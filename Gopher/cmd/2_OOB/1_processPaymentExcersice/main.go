package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
)

// PaymentProcessor interface
type PaymentProcessor interface {
	Pay(amount float64) error
	Refund(amount float64) error
}

// CreditCardProcessor implementation
type CreditCardProcessor struct {
	holderNumber int
	holderName   string
	limit        float64
}

func (c *CreditCardProcessor) Pay(amount float64) error {
	if amount > c.limit {
		return ErrInsufficientBalance
	}
	c.limit -= amount
	fmt.Printf("Paid $%.2f using Credit Card (%s)\n", amount, c.holderName)
	return nil
}

func (c *CreditCardProcessor) Refund(amount float64) error {
	c.limit += amount
	fmt.Printf("Refunded $%.2f to Credit Card (%s)\n", amount, c.holderName)
	return nil
}

// PayPalProcessor implementation
type PayPalProcessor struct {
	email   string
	balance float64
}

func (p *PayPalProcessor) Pay(amount float64) error {
	if amount > p.balance {
		return ErrInsufficientBalance
	}
	p.balance -= amount
	fmt.Printf("Paid $%.2f using PayPal (%s)\n", amount, p.email)
	return nil
}

func (p *PayPalProcessor) Refund(amount float64) error {
	p.balance += amount
	fmt.Printf("Refunded $%.2f to PayPal (%s)\n", amount, p.email)
	return nil
}

// ProcessPayment uses the interface — it doesn’t care *which* processor it is
func ProcessPayment(p PaymentProcessor, amount float64) {
	if err := p.Pay(amount); err != nil {
		switch {
		case errors.Is(err, ErrInsufficientBalance):
			log.Println("Payment failed:", err)
		default:
			log.Println("Unexpected error:", err)
		}
	} else {
		log.Println("Payment processed successfully.")
	}
}

func main() {
	cc := &CreditCardProcessor{holderName: "Saleh", holderNumber: 123456, limit: 100}
	pp := &PayPalProcessor{email: "saleh@example.com", balance: 75}

	ProcessPayment(cc, 50)
	ProcessPayment(pp, 30)
	ProcessPayment(pp, 100) // should trigger insufficient balance
}
