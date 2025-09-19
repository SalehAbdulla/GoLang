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
	if len(args) != 1 {
		log.Fatal("expected 1 argument")
	}
	// input := args[0]

	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}

	asciiMap := utils.GetAscii(asciiArt)

	// for r, arts := range asciiMap {
	// 	fmt.Println(r)
	// 	for _, art := range arts {
	// 		fmt.Println(art)
	// 	}
	// 	fmt.Println(strings.Repeat("-", 30))
	// }

	commands.RenderAscii(args[0], asciiMap, os.Stdout)
}
