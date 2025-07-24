package piscine

// Takes a string of space-separated numbers (like "1 2 3 4 5 26") 
// and returns the corresponding lowercase alphabet letters joined as a single string (like "abcdez").
// Then, reverse the process: take any lowercase word (like "hello") and return the space-separated numbers (like "8 5 12 12 15"),
// where a = 1, z = 26.

func AsciiMagic(input string) string {
	inputToRunes := []rune(input)
	result := ""
	for _, char := range inputToRunes {
		if char == ' ' {
			continue
		}
		// here Ascii Transliation
		if char >= '0' && char <= '9' {
			result += string('a' + (atoi3(string(char)) - 1)) // or -96
		} 
		if char >= 'a' && char <= 'z' {
			result += itoa3(int(char - 96))
			result += " "
		}
	}

	return result + "\n"
}

func atoi3(number string) int {
	// check if empty
	if number == "" {return 0}
	result := 0
	for _, nbChar := range number {
		if nbChar >= '0' && nbChar <= '9' {
			result *= 10
			result += int(nbChar - '0')
		} else {
			return 0
		}
	}
	return result
}

func itoa3(number int) string {
	// if number == 0
	if number == 0 {return "0"}
	// Handle negative
	isNegative := false 
	if number < 0 {
		isNegative = true
		number = -number
	}
	
	result := ""
	for number > 0 {
		digit := number % 10
		result = string(digit + '0') + result
		number /= 10
	}

	if isNegative {
		return "-" + result
	} else {
		return result
	}
}
