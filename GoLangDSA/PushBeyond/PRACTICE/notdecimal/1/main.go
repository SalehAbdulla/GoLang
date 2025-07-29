package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {

	//// If the argument is empty return a newline \n.
	if dec == "" {
		return "\n"
	}

	//// If the number doesn't have a decimal point or there is only a zero after the . return the number followed by a newline \n.

	dotExist := -1 // by default false
	for i, char := range dec {
		if char == '.' {
			dotExist = i
			break
		}
	}

	if dotExist == -1 {return dec + "\n"}

	isAllZerosAfterDot := true

	for i := dotExist+1; i < len(dec); i++ {
		if dec[i] != '0' {
			isAllZerosAfterDot = false
			break
		}
	}

	if isAllZerosAfterDot {return dec + "\n"}

	/// ----
	// If the argument is not a number return it followed by a newline \n.
	cleanedDec := "" // Remove . from the string
	for _, char := range dec {
		if char != '.'{
			cleanedDec += string(char)
		}
	}

	num := Atoi(cleanedDec)
	if num == -1 {return dec + "\n"}

	return Itoa(num) + "\n"

}

func Atoi(num string) int {
	if num == "" {return -1}
	if num == "0" {return 0}

	isNegative := false
	if num[0] == '-' {
		isNegative = true
		num = num[1:]
	} 
	// -----------
	result := 0
	for _, n := range num {
		if n >= '0' && n <= '9' {
			result *= 10
			result += int(n - '0')
		} else {
			return -1
		}
	}
	// -----------
	if isNegative {
		return result *-1
	}
	return result
}

func Itoa(num int) string {
	if num == 0 {return "0"}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	// -----------
	result := ""
	for num > 0 {
		digit := num % 10
		result = string(rune(digit + '0')) + result
		num /= 10
	}
	// -----------

	if isNegative {
		return "-" + result
	}
	return result
}
