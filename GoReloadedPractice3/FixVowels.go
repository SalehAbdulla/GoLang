package reloaded

import "strings"

func FixVowels(s string) string {
	if strings.TrimSpace(s) == "" {
		return s
	}
	toSlice := strings.Fields(s)
	for i := 0; i < len(toSlice) -1; i++{
		next := []rune(toSlice[i+1])
		if strings.ToLower(toSlice[i]) == "a" {
			if IsVowel(next[0]) {
				if toSlice[i] == "A" {
					toSlice[i] = "An"
				} else {
					toSlice[i] = "an"
				}
			}
		}
	}
	s = strings.Join(toSlice, " ")
	return s
}

func IsVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
