package main

import (
	"log"
	"os"
	"reloaded"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Arguments must be 2")
	}

	fileNameToMod := args[0]

	data, err := os.ReadFile(fileNameToMod)

	if err != nil {
		log.Fatal(err.Error())
	}

	result := reloaded.ProcessCommands(string(data))

	err2 := os.WriteFile(args[1], []byte(result), 0644)
	if err2 != nil {log.Fatal(err2.Error())}
}
