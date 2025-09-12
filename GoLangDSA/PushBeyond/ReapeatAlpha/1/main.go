package main

import "fmt"

func RepeatAlpha(s string) string {
	var result string
	for _, char := range s {
		getCount := int(char - 96) // a -> 1 // b -> 2
		for i := 0; i < getCount; i++ {
			result += string(char)
		}
	}
	return result
}

func main() {
	text := "abcd"
	fmt.Println(RepeatAlpha(text))
}
