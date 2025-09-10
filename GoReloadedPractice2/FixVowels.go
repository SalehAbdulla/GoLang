package reloaded

import "strings"

func FixVowels(ref string) string {
	if strings.TrimSpace(ref) == "" {return ref}
	words := strings.Fields(ref)

	for i := 0; i < len(words) -1; i++ {
		next := []rune(words[i+1])
		if strings.ToLower(words[i]) == "a" {
			if isVowel(next[0]) {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}
	ref = strings.Join(words, " ")
	return ref
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'u' || r == 'i' || r == 'o' || r == 'A' || r == 'E' || r == 'U' || r == 'I' || r == 'O'
}