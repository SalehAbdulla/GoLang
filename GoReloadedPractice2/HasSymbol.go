package reloaded

import "unicode"

func HasSymbol(text string) bool {
	if text == "" {
		return false
	}

	for _, char := range text {
		if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			return true
		}
	}
	
	return false
}