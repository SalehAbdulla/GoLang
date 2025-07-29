package piscine

// Write a function `WordFlip()` that takes a `string` as input and returns it in reverse order.

// - The output should be followed by a newline `\n`.

// - If the string is empty, return `Invalid Output`.
// - Ignore multiple spaces between words and trim any leading or trailing spaces in the string.

	
func WordFlip(str string) string {
	if str == "" {return "Invalid Output"}
	words := []string{}
	word := ""

	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}

	if word != "" {
		words = append(words, word)
	}

	result := ""
	for i := len(words) -1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}

	return result + "\n"
}
