package piscine

func DigitLen(n, base int) int {
	// The second int must be between 2 and 36. If not, the function returns -1.
	if !(base >= 2 && base <= 36) {
		return -1
	// If the first int is negative, reverse the sign and count the digits.
	}
	if n < 0 {
		n = -n
	}
	// returns the times the first int can be divided by the second until it reaches zero.
	counter := 0
	for n > 0 {
		n = n / base
		counter++
	}
	return counter
}
