package piscine

func SplitWhiteSpaces(s string) []string {
	sToRunes := []rune(s)
	result := []string{}
	word := ""

	for i, char := range sToRunes {
		if !(char == ' ' || char == '\t' || char == '\n') {
			word += string(rune(char))
		} else if i == len(sToRunes)-1 {
			word += string(rune(char))
		} else {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		}
	}

	lastChar := sToRunes[len(sToRunes)-1]
	lastWordResult := result[len(result)-1]
	newWord := lastWordResult + string(lastChar)
	result[len(result)-1] = newWord

	return result
}
