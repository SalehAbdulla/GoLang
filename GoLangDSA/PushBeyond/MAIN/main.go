package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func applyMultAdd(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {

	arrayFunction := []func(int, int) int{add, multiply}
	
	sum := applyMultAdd(5, 3, arrayFunction[0])
	product := applyMultAdd(5, 3, arrayFunction[1])

	fmt.Println("The value of the sum, ", sum)
	fmt.Println("The value of the product, ", product)

}
