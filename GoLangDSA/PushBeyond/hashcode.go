package piscine

// (ASCII of current character + size of the string) % 127, ensuring the result falls within the ASCII range of 0 to 127.
// If the resulting character is unprintable add 33 to it.

func HashCode(dec string) string {
	result := ""
	for _, char := range dec {
		encodedChar := (int(char) + len(dec) ) %127
		if encodedChar < 33 {
			encodedChar += 33
		}
		result += string(encodedChar)
	}
	return result
}
