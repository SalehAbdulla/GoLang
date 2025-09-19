package utils

import "asciiArt/internal/constants"

func GetAsciiMap(asciiMapSlice []string) map[rune][]string {
	hashMap := make(map[rune][]string)
	
	currentRune := constants.START_RUNE
	var buffer []string

	for _, val := range asciiMapSlice {
		if val == "" && len(buffer) > 0 {
			hashMap[rune(currentRune)] = buffer
			currentRune++
			buffer = nil
		} else {
			buffer = append(buffer, val)
		}
	}

	if len(buffer) > 0 {
		hashMap[rune(currentRune)] = buffer
	}
	
	return hashMap
}