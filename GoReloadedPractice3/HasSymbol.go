package reloaded

import "unicode"

func HasSymbol(s string) bool {
	for _, char := range s {
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			return true
		}
	}
	return false
}