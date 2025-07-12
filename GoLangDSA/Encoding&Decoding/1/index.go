package main

import (
	"fmt"
	"strconv"
)

func main(){

}


func encode(strs []string) string {
	result := ""
	for _, word := range strs {
		// len(str) + "#" + str 5#Hello
		encoded := strconv.Itoa(len(word)) + "#" + word
		result += encoded
	} 
	return result
}

func decode(encoded string) []string {
	result := []string{}

	i := 0

	for i < len(encoded) {
		j := i
		
		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])
		
		i = j + 1
		result = append(result, encoded[i:i + length])
		i += length

	}

	return result

}


