package main

import (
	"asciiArt/internal/commands"
	"asciiArt/internal/utils"
	"bufio"
	"log"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) != 1 {log.Fatal("Expected 1 argument")}
	
	input := args[0]

	var asciiArtToSlice []string
	file, err := os.Open("standard.txt")
	if err != nil {log.Fatal(err.Error())}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		asciiArtToSlice = append(asciiArtToSlice, scanner.Text())
	}

	getAsciiMap := utils.GetAsciiMap(asciiArtToSlice)
	commands.RenderAscii(input, getAsciiMap, os.Stdout)
}
