package main

import (
	"log"
	"os"
	"strings"
	"unicode"
	reloaded "reloaded/internal/commands"
)

func main() {
	// First I will take the args, make sure we got 2 and ends with .txt
	args := os.Args[1:]
	if len(args) != 2 {log.Fatal("Arguments must be 2.")}
	if !strings.HasSuffix(args[0], ".txt") || !strings.HasSuffix(args[1], ".txt") {log.Fatal("Arguments must end with .txt")}

	// Second that I've took the args, and make sure that they are text files,
	// I'm going to read the file, make sure the file not exceeded 1000 byte, and only ASCII is supported.
	fileData, err := os.ReadFile(args[0])
	if err != nil {log.Fatal(err.Error())}

	// Third thing and better to be at the beginning, making sure the file is not empty
	// we don't want to consume performace; so, we will return it back from the beginning.
	if strings.TrimSpace(string(fileData)) == "" {
		log.Fatal("File is empty or contains white spaces only.")
	}

	const maxData = 1000
	if len(fileData) > maxData {log.Fatal("file data exceeds the limit - 1000 byte is supported.")}

	for _, r := range string(fileData) {
		if r > unicode.MaxASCII {
			log.Fatalf("Unspported character - CLI app supports ASCII-only. char-unspported: %q", r)
		}
	}

	// Then that we've made sure we've got a supported text, we will proceed with the next function
	// processCommands

	getResult := reloaded.ProcessCommands(string(fileData))
	
	// after processing, I'm going to write the result back to result.txt

	err = os.WriteFile(args[1], []byte(getResult), 0664)
	if err != nil {log.Fatal(err.Error())}
}
