package utils

import "asciiArt/internal/constants"

func GetAscii(standard []string) map[rune][]string {
	asciiMap := make(map[rune][]string)
	currentRune := constants.START_RUNE

	var buffer []string

	for _, str := range standard {
		if str == "" && len(buffer) > 0 {
			asciiMap[rune(currentRune)] = buffer
			currentRune++
			buffer = nil
		} else if str != "" {
			buffer = append(buffer, str)
		}
	}

	if len(buffer) > 0 {
		asciiMap[rune(currentRune)] = buffer
	}

	return asciiMap
}