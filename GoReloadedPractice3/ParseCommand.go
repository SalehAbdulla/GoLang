package reloaded

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func ParseCommand(currentCommand string) (string, int) {
	cmd := ""
	count := 1
	for _, char := range currentCommand {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			cmd += string(char)
		}
	}

	if containsDigits(currentCommand) {
		getCount := ""
		for _, char := range currentCommand {
			if unicode.IsDigit(char) {
				getCount += string(char)
			}
		}
		if getCount != "" {
			toInt, err := strconv.Atoi(getCount)
			if err != nil {log.Fatal(err.Error())}
			count = toInt
		}
	}

	return strings.ToLower(cmd), count
}

func containsDigits(s string) bool {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
