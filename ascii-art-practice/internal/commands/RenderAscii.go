package commands

import (
	"asciiArt/internal/constants"
	"fmt"
	"io"
	"log"
	"strings"
)

func RenderAscii(input string, asciiArt map[rune][]string, w io.Writer) {
	input = strings.ReplaceAll(input, `\n`, "\n")
	if input == "\n" {
		fmt.Fprintln(w); return
	}

	inputLines := strings.Split(input, "\n")

	for _, inputLine := range inputLines {
		// this is because the split function above could results an empty str []"Hello","","World"]

		// if input is already empty
		// if there's consecutive \n\n\n\n\n
		// if the input ends with \n -> hello\n

		if inputLine == "" {
			fmt.Fprintln(w); continue
		}

		var outputBuilders [constants.RUNE_HEIGHT]strings.Builder

		for _, ch := range inputLine {
			if art, ok := asciiArt[ch]; ok {
				for i := 0; i < constants.RUNE_HEIGHT; i++ {
					outputBuilders[i].WriteString(art[i])
				}
			} else {
				log.Fatalf("unsupported character %q", ch)
			}
		}

		for _, out := range outputBuilders {
			fmt.Fprintln(w, out.String())
		}

	}


}
