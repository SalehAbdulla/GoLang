package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	var asciiArtSlice []string
	file, err := os.Open("standard.txt")
	if err != nil {log.Fatal(err.Error())}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {asciiArtSlice = append(asciiArtSlice, scanner.Text())}

	for _, v := range asciiArtSlice {
		fmt.Println(v)
	}

	
	
}
