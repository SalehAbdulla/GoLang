package reloaded

import "strings"

func ReloadVowels(s string) string {
	strs := strings.Fields(s)
	for i := 0; i < len(strs)-1; i++ {
		next := []rune(strs[i+1])
		if strings.ToLower(strs[i]) == "a" {
			if isVowel(next[0]) {
				if strs[i] == "A" {
					strs[i] = "An"
				} else {
					strs[i] = "an"
				}
			}
		}
	}
	s = strings.Join(strs, " ")
	return s
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'i' || r == 'e' || r == 'u' || r == 'o' || r == 'A' || r == 'I' || r == 'E' || r == 'U' || r == 'O'
}