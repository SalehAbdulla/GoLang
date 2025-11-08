package main

import "fmt"

// Inheritance, “Reuse through hierarchy”

type Vehicle struct{}

func (v Vehicle) Start() {
	fmt.Println("Starting vehicle...")
}

type Car struct {
	Vehicle // embedding Inheritance
}

func (c Car) Drive() {
	fmt.Println("Driving car...")
}

func main() {
	c := Car{}
	c.Start() // works! thanks to embedding
	c.Drive()
}
