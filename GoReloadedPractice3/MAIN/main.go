package main

import (
	"log"
	"os"
	"reloaded"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Arguments must be 2")
	}

	if !strings.HasSuffix(args[0], ".txt") && !strings.HasSuffix(args[1], ".txt") {
		log.Fatal("Arguments names must end with .txt")
	}

	data, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err.Error())
	}

	result := reloaded.ProcessCommands(string(data))

	err = os.WriteFile(args[1], []byte(result), 0644)
	if err != nil {log.Fatal(err.Error())}
}
