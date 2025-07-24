package piscine

func itoa(number int) string {
	// Handle zero
	if number == 0 {
		return "0"
	}
	// handle negative sign
	isNegative := false
	if number < 0 {
		isNegative = true
		number = -number
	}
	// Algo
	result := ""
	for number > 0 {
		digit := number % 10
		result = string(rune(digit+'0')) + result
		number /= 10
	}
	// return
	if isNegative {
		return "-" + result
	} else {
		return result
	}
}
