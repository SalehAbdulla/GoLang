package piscine

func itoa(number int) string {
	if number == 0 {
		return "0"
	}

	result := ""
	isNegative := false
	if number < 0 {
		isNegative = true
		number = -number
	}

	for number > 0 {
		digit := number % 10
		result = string(rune(digit+'0')) + result
		number /= 10
	}

	if isNegative {
		return "-" + result
	} else {
		return result
	}
}
