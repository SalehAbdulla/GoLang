package piscine

// String to Integer
func atoi(number string) int {
	// Handle Nothing
	if number == "" {
		return 0
	}
	// Check if number is negative
	isNegative := false
	if number[0] == '-' {
		isNegative = true
		number = number[1:] // get rid of - sign
	}
	result := 0
	for _, digit := range number {
		// loop through number - check if valid
		if digit >= '0' && digit <= '9' {
			result *= 10
			result += int(digit - '0') // Convert each digit into actual int(num - '0')
		} else {
			return 0
		}
	}
	if isNegative {
		return -result
	} else {
		return result
	}
}
