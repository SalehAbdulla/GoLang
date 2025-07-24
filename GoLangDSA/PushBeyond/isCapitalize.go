package piscine

func IsCapitalized(s string) bool {
	if s == "" {return false}

	newWord := true // Keep track of new word
	for _, r := range s {
		if r == ' ' {
			newWord = true
		} else {
			if newWord {
				// Check if the first character of a word is lowercase
				if r >= 'a' && r <= 'z' {
					return false
				}
				newWord = false
			}
		}
	}
    
	return true
}
