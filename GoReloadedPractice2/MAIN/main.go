package main

import (
	"log"
	"os"
	"reloaded"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {log.Fatal("arguments passed must be 2.")}

	if !strings.HasSuffix(args[0], ".txt") || !strings.HasSuffix(args[1], ".txt") {
		log.Fatal("arguments type must be .txt")
	}

	data, err := os.ReadFile(args[0])
	if err != nil {log.Fatal(err.Error())}

	result := reloaded.ProcessCommands(string(data))

	err2 := os.WriteFile(args[1], []byte(result), 0644)
	if err2 != nil {log.Fatal(err2.Error())}
}
