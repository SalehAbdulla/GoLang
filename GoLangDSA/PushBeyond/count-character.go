package piscine

///Instructions
// write a function that takes a string and a character as arguments and returns the number of times the character appears in the string.

// if the character is not in the string return 0
// if the string is empty return 0

func CountChar2(str string, c rune) int {
	if str == "" {return 0}
	counter := 0
	for _, char := range str {
		if char == c {
			counter++
		}
	}
	return counter
}
