package commands

import (
	"asciiArt/internal/constants"
	"fmt"
	"io"
	"log"
	"strings"
)

func RenderAscii(input string, asciiMap map[rune][]string, w io.Writer) {
	if input == "" {
		return
	}
	input = strings.ReplaceAll(input, `\n`, "\n")

	onlyNewLines := true
	for _, char := range input {
		if char != '\n' {
			onlyNewLines = false
			break
		}
	}

	if onlyNewLines {
		countNewLines := strings.Count(input, "\n")
		for countNewLines > 0 {
			fmt.Fprintln(w)
			countNewLines--
		}
		return
	}

	splitInputByNewLine := strings.Split(input, "\n")

	for _, line := range splitInputByNewLine {
		var buffer [constants.RUNE_HEIGHT]strings.Builder

		for _, char := range line {
			if asciiArt, ok := asciiMap[char]; ok {
				for i := 0; i < constants.RUNE_HEIGHT; i++ {
					buffer[i].WriteString(asciiArt[i])
				}
			} else {
				log.Fatal("unsupported character", char)
			}
		}

		for i := 0; i < len(buffer); i++ {
			fmt.Fprintln(w, buffer[i].String())
		}

	}

}
