package piscine


func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	const digits = "0123456789ABCDEF"
	var result string
	var uvalue uint64

	if value == 0 {
		return "0"
	}

	isNegative := value < 0
	if isNegative {
		uvalue = uint64(-int64(value)) // safely convert to uint64
	} else {
		uvalue = uint64(value)
	}

	for uvalue > 0 {
		remainder := uvalue % uint64(base)
		result = string(digits[remainder]) + result
		uvalue /= uint64(base)
	}

	if isNegative {
		result = "-" + result
	}

	return result
}
