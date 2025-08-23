package piscine

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}
	digits := "0123456789ABCDEF"
	var uValue uint64
	var ubase uint64
	isNegative := value < 0
	
	if isNegative {
		uValue = uint64(-int64(value))
	} else {
		uValue = uint64(value)
	}

	ubase = uint64(base)

	var result string

	for uValue > 0 {
		getReminder := uValue % ubase
		result = string(digits[getReminder]) + result
		uValue /= ubase
	}

	if isNegative {
		return "-" + result
	}
	return result
}
