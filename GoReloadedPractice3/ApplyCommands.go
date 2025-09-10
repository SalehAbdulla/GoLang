package reloaded

import (
	"log"
	"strconv"
	"strings"
)

func ApplyCommands(ref, cmd string, count, commandStarts int) string {
	before := strings.TrimSpace(ref[:commandStarts])
	

	words := strings.Fields(before)
	
	if len(words) == 0 {
		log.Fatal("No Words Before the command")
	}

	// Edge Case -> four words to cap (cap, 10)
	if count > len(words) {
		count = len(words)
	}

	start := len(words) - count

	for i := start; i < len(words); i++ {
		switch cmd {
			case "low":
				words[i] = strings.ToLower(words[i])
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "cap":
				wordToRune := []rune(words[i])
				if len(wordToRune) > 0 {
					words[i] = strings.ToUpper(string(wordToRune[:1])) + strings.ToLower(string(wordToRune[1:]))
				}
			case "hex":
				if val, err := strconv.ParseInt(words[i], 16, 64); err == nil {
					words[i] = strconv.FormatInt(val, 10)
				}
			case "bin":
				if val, err := strconv.ParseInt(words[i], 2, 64); err == nil {
					words[i] = strconv.FormatInt(val, 10)
				} 
			}
	}

	wordsBeforeToStr := strings.Join(words, " ")

	return wordsBeforeToStr + ref[commandStarts:]
}


