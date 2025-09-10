package main

import (
	"os"
	"reloaded"
	"strings"
	"log"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {log.Fatal("Expected 2 arguments")}

	if !strings.HasSuffix(args[0], ".txt") || !strings.HasSuffix(args[1], ".txt") {
		log.Fatal("Text Files must end with .txt")
	}

	userInput, err := os.ReadFile(args[0])
	if err != nil {log.Fatal("Error reading file:", err.Error())}

	textResult := reloaded.ProcessCommands(string(userInput))

	err = os.WriteFile(args[1], []byte(textResult), 0666)
	if err != nil {log.Fatal("Error writing file:", err.Error())}
}
