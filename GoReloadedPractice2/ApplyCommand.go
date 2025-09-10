package reloaded

import (
	"log"
	"strconv"
	"strings"
)

func ApplyCommand(ref, cmd string, count, cmdIndex int) string {
	before := strings.TrimSpace(ref[:cmdIndex])
	words := strings.Fields(before)
	after := ref[cmdIndex:]

	if len(words) == 0 {log.Fatal("No Words Before The Command")}
	if count > len(words) {count = len(words)}

	start := len(words) - count
	for i := start; i < len(words); i++ {
		switch cmd {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			if len(words[i]) > 0 {
				toSlice := []rune(words[i])
				words[i] = strings.ToUpper(string(toSlice[:1])) + strings.ToLower(string(toSlice[1:]))
			}
		case "hex":
			num, err := strconv.ParseInt(words[i], 16, 64)
			if err != nil {log.Fatal(err.Error())}
			words[i] = strconv.FormatInt(num, 10)
		case "bin":
			num, err := strconv.ParseInt(words[i], 2, 64)
			if err != nil {log.Fatal(err.Error())}
			words[i] = strconv.FormatInt(num, 10)
		}
	}

	before = strings.Join(words, " ")
	after = removeCommand(after)

	return after + " " + before
}

func removeCommand(text string) string {
	from := strings.Index(text, "(")
	to := strings.Index(text, ")")
	if from != -1 && to != -1 && to > from {
		text = text[:from] + text[to+1:]
	}
	return text
}