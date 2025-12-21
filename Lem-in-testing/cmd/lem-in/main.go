package main

import (
	"lem-in/internal"
	"lem-in/internal/utils"
	"log"
	"os"
	"strings"
)

func main() {
	// 1) Read the file, handle edge cases
	args := os.Args[1:]
	if len(args) != 1 {log.Fatal("error usage: go run . <file.txt>")}
	if !strings.HasSuffix(args[0], ".txt") {log.Fatal("error usage: extention must end with .txt go run . <file.txt>")}

	file, err := os.ReadFile(args[0])
	internal.Check(err)
	println(string(file))
	// 2) Filtering - Remove Comments

	splitFile := strings.Split(utils.RemoveComments(string(file)), "\n")
	for _, val := range splitFile {println(val)}
	






}
