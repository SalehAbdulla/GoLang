package piscine

func atoi(number string) int {
	result := 0
	isNegative := false
	for i, num := range number {
		if i == 0 && num == '-' {
			isNegative = true
			continue
		}
		if num >= '0' && num <= '9' {
			result *= 10
			result += int(num - '0') // int(num- '0') to actual number
		} else {
			return 0
		}
	}
	if isNegative {
		return result * -1
	} else {
		return result
	}
}
