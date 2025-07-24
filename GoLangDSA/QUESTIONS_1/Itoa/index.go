package main

import (
	"fmt"
)

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}

//Write a function that simulates the behavior of the Itoa function in Go. Itoa transforms a number represented as an int in a number represented as a string.

// For this exercise the handling of the signs + or - does have to be taken into account.

func Itoa(n int) string {
	
	// Conver the number into a byte(number + '0') as ascii
	// check if negative

	if n == 0 {
		return "0" 
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n // get rid of - sign n = n*-1 same as n = -n
	}

	nToByte := []byte{}
	for n > 0 {
		nToByte = append(nToByte, byte('0' + n % 10))
		n = n / 10
	}

	// Swap then return
	for i, j := 0, len(nToByte) - 1; i < j; i, j = i + 1, j - 1 {
		nToByte[i], nToByte[j] = nToByte[j], nToByte[i]
	} 

	if isNegative {
		return "-" + string(nToByte)
	} else {
		return string(nToByte)
	}

}

