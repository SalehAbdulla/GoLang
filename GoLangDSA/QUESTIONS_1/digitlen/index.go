package main

import (
	"fmt"
)

func main() {
	fmt.Println(DigitLen(100, 10))
	fmt.Println(DigitLen(100, 2))
	fmt.Println(DigitLen(-100, 16))
	fmt.Println(DigitLen(100, -1))
}

//Write a function DigitLen() that takes two integers as arguments and returns the times the first int can be divided by the second until it reaches zero.

// The second int must be between 2 and 36. If not, the function returns -1.
// If the first int is negative, reverse the sign and count the digits.

func DigitLen(n, base int) int {
	// The second int must be between 2 and 36. If not, the function returns -1.
	if !(base >= 2 && base <= 36) {
		return -1
	} 

	// If the first int is negative, reverse the sign and count the digits.
	if n < 0 {
		n = n * -1
	}

	// returns the times the first int can be divided by the second until it reaches zero
	times := 0

	for n != 0 {
		n /= base
		times++
	}

	return times
}


//$ go run . | cat -e
// 3$
// 7$
// 2$
// -1$