package piscine

func ItoaBase(value, base int) (result string) {
	digits := "0123456789ABCDEF"
	isNegative := value < 0
	var uValue uint64

	if isNegative {
		uValue = uint64(-int64(value))
	} else {
		uValue = uint64(value)
	}

	uBase := uint64(base)

	for uValue > 0 {
		reminder := uValue % uBase
		result = string(digits[reminder]) + result
		uValue /= uBase
	}

	if isNegative {
		result = "-" + result
	}

	return
}
