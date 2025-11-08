package main

import "fmt"

// Polymorphism — “One interface, many forms”

type Shape interface {
	Draw()
}

type Circle struct{}
type Square struct{}

func (Circle) Draw() { fmt.Println("Drawing circle") }
func (Square) Draw() { fmt.Println("Drawing square") }

func DrawShape(s Shape) {
	s.Draw()
}

func main() {
	DrawShape(Circle{})
	DrawShape(Square{})
}
