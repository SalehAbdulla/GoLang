package main

//$ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
// 1$
import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	args := os.Args[1:]

	word1 := args[0]
	word2 := args[1]

	i := 0
	for _, char := range word2 {
		if i < len(word1) && char == rune(word1[i]) {i++}
	}


	if i == len(word1) {
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
