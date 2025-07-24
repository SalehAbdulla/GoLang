package main

import "fmt"

func main(){
	fmt.Println(factorial(10))
}

func factorial(nb int) int { 
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= nb; i++ {
		result = result * i
	}

	return result
}
