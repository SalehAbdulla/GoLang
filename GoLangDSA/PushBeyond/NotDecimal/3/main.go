package piscine

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	decimalIndex := -1
	for i, char := range dec {
		if char == '.' {
			decimalIndex = i
			break
		}
	}
	if decimalIndex == -1 {
		return dec + "\n"
	}
	allZerosAfterDot := true
	for i := decimalIndex + 1; i < len(dec); i++ {
		if dec[i] != '0' {
			allZerosAfterDot = false
			break
		}
	}
	if allZerosAfterDot {
		return dec + "\n"
	}
	var dotCleared string
	for _, char := range dec {
		if char != '.' {
			dotCleared += string(char)
		}
	}
	inputToInt, ok := Atoi(dotCleared)
	if !ok {
		return dec + "\n"
	}
	return Itoa(inputToInt) + "\n"
}

func Atoi(num string) (int, bool) {
	if num == "0" {
		return 0, true
	}
	isNegative := false
	if num[0] == '-' {
		isNegative = true
		num = num[1:]
	}
	var result int
	for _, char := range num {
		if char < '0' || char > '9' {
			return -1, false
		} else {
			result *= 10
			result += int(char - '0')
		}
	}
	if isNegative {
		return result * -1, true
	}
	return result, true
}

func Itoa(num int) string {
	if num == 0 {
		return "0"
	}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	var result string
	for num > 0 {
		digit := num % 10
		result = string(rune(digit+'0')) + result
		num /= 10
	}
	if isNegative {
		return "-" + result
	}
	return result
}
