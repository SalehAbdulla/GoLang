package main

import (
	"os"
	"strings"
)

func main(){
	args := os.Args[1:]
	if len(args) != 2 {return}

	text, _ := os.ReadFile(args[0])
	os.WriteFile(args[1], []byte(FixVowels(string(text))), 0644)
}


func IsVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'u' || r == 'o' || r == 'i' || r == 'A' || r == 'E' || r == 'U' || r == 'O' || r == 'I'
}

// FixVowels Logic ---------------------

func FixVowels(s string) string {
	lines := strings.Split(s, "\n")
	for li, line := range lines {
		words := strings.Fields(line)
		for i := 0; i < len(words)-1; i++ {
			next := []rune(words[i+1])
			if strings.ToLower(words[i]) == "a" {
				if len(next) > 0 && IsVowel(next[0]) {
					if words[i] == "A" {
						words[i] = "An"
					} else {
						words[i] = "an"
					}
				} 
			}
		}
		lines[li] = strings.Join(words, " ")
	}
	return strings.Join(lines, "\n")
}

