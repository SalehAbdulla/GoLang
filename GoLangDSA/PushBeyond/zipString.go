package piscine

import "strconv"

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	runes := []rune(s)
	result := ""
	count := 1

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			count++
		} else {
			result += strconv.Itoa(count) + string(runes[i-1])
			count = 1
		}
	}
    
	result += strconv.Itoa(count) + string(runes[len(runes)-1]) // lastChar will be missed
	return result
}
