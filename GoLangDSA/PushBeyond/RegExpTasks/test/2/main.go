package main

import (
	"fmt"
	"regexp"
)

func main() {
	tests := []string{
		"((((((((up))))))))",
		"(((hex)))",
		"((((up, 2))))",
		"((((((bin,10))))))",
	}

	re := regexp.MustCompile(`\(+\s*(hex|bin|low|up|cap)(\s*,\s*\d+)?\s*\)+`)

	for _, t := range tests {
		result := re.ReplaceAllString(t, "($1$2)")
		fmt.Println(result)
	}
}
