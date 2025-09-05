package main

import (
	"fmt"
	"strings"
)

func TrimSingleQuotes(s string) string {
	result := ""
	inQuote := false
	segment := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '\'' {
			if inQuote {
				// closing single quote
				inQuote = false
				result += "'" + strings.TrimSpace(segment) + "'"
				segment = ""
			} else {
				// opening single quote
				inQuote = true
			}
		} else if inQuote {
			segment += string(char)
		} else {
			result += string(char)
		}
	}
	// append any unclosed segment
	if segment != "" {
		result += "'" + strings.TrimSpace(segment)
	}
	return result
}

func TrimDoubleQuotes(s string) string {
	result := ""
	inQuote := false
	segment := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '"' {
			if inQuote {
				// closing double quote
				inQuote = false
				result += `"` + strings.TrimSpace(segment) + `"`
				segment = ""
			} else {
				// opening double quote
				inQuote = true
			}
		} else if inQuote {
			segment += string(char)
		} else {
			result += string(char)
		}
	}
	// append any unclosed segment
	if segment != "" {
		result += `"` + strings.TrimSpace(segment)
	}
	return result
}

// Combine both
func TrimQuotes(s string) string {
	// First trim single quotes
	s = TrimSingleQuotes(s)
	// Then trim double quotes
	s = TrimDoubleQuotes(s)
	return s
}

func main() {
	fmt.Println(TrimQuotes(`"' world '"`))            // -> "'world'"
	fmt.Println(TrimQuotes(`' hello '`))              // -> "'hello'"
	fmt.Println(TrimQuotes(`"example"`))              // -> '"example"'
	fmt.Println(TrimQuotes(`It is ' cool ' indeed`))  // -> "It is 'cool' indeed"
	fmt.Println(TrimQuotes(`"   'nested quote'   "`)) // -> "\"'nested quote'\""
	fmt.Println(TrimQuotes(`He said ' hello ' and she said " world "`)) // -> He said 'hello' and she said "world"
}
