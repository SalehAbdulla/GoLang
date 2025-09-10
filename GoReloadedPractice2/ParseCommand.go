package reloaded

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

func ParseCommand(cmd string) (string, int) {
	command := ""
	count := 1

	for _, char := range cmd {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			command += string(char)
		}
	}

	if containsDigits(cmd) {
		getNum := ""
		for _, char := range cmd {
			if unicode.IsDigit(char) {
				getNum += string(char)
			}
		}
		if len(getNum) > 0 {
			toInt, err := strconv.Atoi(getNum)
			if err != nil {log.Fatal(err.Error())}
			count = toInt
		}
	}
	
	return strings.ToLower(command), count
}

func containsDigits(cmd string) bool {
	for _, char := range cmd {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}