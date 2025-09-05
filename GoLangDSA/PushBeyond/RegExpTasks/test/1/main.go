package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//' apple                           ' g '
	// 'apple' g'

	s := `' apple                           ' g '`

	i := 0
	for {
		//' apple                           ' g '

		start := strings.Index(s[i:], "'") // 0
		if start == -1 {
			break
		}
		start += i

		if (start+1 < len(s) && unicode.IsLetter(rune(s[start+1]))) &&
			(start-1 >= 0 && unicode.IsLetter(rune(s[start-1]))) {
			i = start + 1
			continue
		}

		end := strings.Index(s[start+1:], "'")
		if end == -1 {
			break
		}
		end += start + 1

		between := strings.TrimSpace(s[start+1 : end])

		s = s[:start] + "'" + between + "'" + s[end+1:]

		i = start + len(between) + 2
	}
	fmt.Println(s)
}
