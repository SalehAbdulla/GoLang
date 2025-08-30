package reloaded

import (
	"log"
	"strconv"
	"strings"
)

func FetchCommand(currentCommand string) (string, int) {
	getCommand := ""
	getWordCount := 1

	containsDigits := func(s string) bool {
		for _, char := range s {
			if char >= '0' && char <= '9' {
				return true
			}
		}
		return false
	}

	if containsDigits(currentCommand) {
		var number string
		for _, char := range currentCommand {
			if containsDigits(string(char)) {
				number += string(char)
			}
		}
		numberToInt, err := strconv.Atoi(number)
		if err != nil {log.Fatal(err.Error())}
		getWordCount = numberToInt
	} 
	
	for _, char := range currentCommand {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' {
			getCommand += string(char)
		}
	}

	return strings.ToLower(getCommand), getWordCount
}
