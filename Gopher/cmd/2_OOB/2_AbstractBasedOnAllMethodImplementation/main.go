package main

import "fmt"

// Abstraction
type Speaker interface {
	Speak() error
}

type Human struct{}

func (h *Human) Speak() error {
	fmt.Println("Hello Everyone!")
	return nil
}

type Animal struct{}

func (a *Animal) Speak() error {
	fmt.Println("beep boop")
	return nil
}
