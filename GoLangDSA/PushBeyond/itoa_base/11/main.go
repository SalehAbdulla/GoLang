package piscine

const DIGITS = "0123456789ABCDEF"

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}

	var uValue uint64
	isNegative := value < 0

	if isNegative {
		uValue = uint64(-int64(value))
	} else {
		uValue = uint64(value)
	}

	var uBase uint64
	uBase = uint64(base)

	var result string

	for uValue > 0 {
		reminder := uValue % uBase
		result = string(DIGITS[reminder]) + result
		uValue /= uBase
	}

	if isNegative {
		return "-" + result
	}

	return result
}
