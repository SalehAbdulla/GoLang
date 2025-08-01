package main

import "fmt"

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	for i := nb - 1; i >= 2; i-- {
		if IsPrime(i) {
			return i
		}
	}
	return -1
}

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	i := 2
	for i*i <= num {
		if num%i == 0 {
			return false
		}
		i++
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(5))
	
}