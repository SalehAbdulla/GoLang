package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "                     it was the worst of times (up) , it was the worst of times (Cap)"
	cmdIndex := strings.Index(text, "(")
	text, _ = ApplyCommand(text, "up", 2, cmdIndex)
	fmt.Println(text)
}

func ApplyCommand(text, cmd string, count, cmdIndex int) (string, error) {
	before := strings.TrimSpace(text[:cmdIndex])
	after := text[cmdIndex:]

	splitBefore := strings.Fields(before)

	start := len(splitBefore) - count
	for i := start; i < len(splitBefore); i++ {
		switch cmd {
			case "up":
				splitBefore[i] = strings.ToUpper(splitBefore[i])
		}
	}
	before = strings.Join(splitBefore, " ")
	return before + " " + removeFirstCommand(after), nil
}

func removeFirstCommand(text string) string {
	from := strings.Index(text, "(")
	to := strings.Index(text, ")")
	if from != -1 && to != -1 && to > from {
		text = strings.TrimSpace(text[:from] + text[to+1:])
	}
	return text
}