package main

import (
	"fmt"
	"regexp"
)

func main(){
	// The Task is to replace this crazy space, to only one whitespace.
	text := "I was thinking ,                     you were right"
	removeSpacesBeforePunc := `\s+([,.!?;:])`
	AddSpaceAfterPunc := `([,.!?;:])([^\s,.!?;:])`
	collapseSpaces := `\s+`
	text = MustCompile(text, removeSpacesBeforePunc, `$1`)
	text = MustCompile(text, AddSpaceAfterPunc, `$1 $2`)
	text = MustCompile(text, collapseSpaces, `$1 `)

	fmt.Println(text)
}

func MustCompile(text, cmd, repl string) string {
	re := regexp.MustCompile(cmd)
	return re.ReplaceAllString(text, repl)
}