package piscine

//Write a function CountAlpha() that takes a string as an argument and returns the number of alphabetic characters in the string.

func CountAlpha(s string) int {
	counter := 0
	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			counter++
		}
	}
	return counter
}
