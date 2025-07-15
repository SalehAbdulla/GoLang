package main

import (
	"fmt"
	"strconv"
)

func main(){
	orginal := []string {"neet","code","loves","you"}
	encoded := encoding(orginal)
	fmt.Println(encoded)
	decoded := decode(encoded) 
	fmt.Println(decoded)
}

func encoding(strs []string) string {
	// Turn it into 4#word4#word
	result := ""
	for _, word := range strs {
		encodedWord := strconv.Itoa(len(word)) + "#" + word
		result += encodedWord
	}
	return result
}

func decode(encoded string) []string {
	result := []string{}

	leftPointer := 0 // i pointer
	for leftPointer < len(encoded) {

		rightPointer := leftPointer

		for encoded[rightPointer] != '#' {
			rightPointer++
		}

		// extract the number from the word   from 0 index till #
		wordSize, _ := strconv.Atoi(encoded[leftPointer:rightPointer])
		leftPointer = rightPointer + 1 // skips # Symbol

		result = append(result, encoded[leftPointer:leftPointer+wordSize])
		leftPointer += wordSize
		fmt.Println(result)

	}

	return result
}