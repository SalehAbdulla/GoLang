package main

import (
	"regexp"
	"fmt"
)

func main(){
	text := "Hello \t     \r\n   , World!"
	cmd := `\s+([.,!?;:])`
	repl := "$1"
	text = MustComplie(text, cmd, repl)
	fmt.Println(text)
}

func MustComplie(s, command, replacement string) string {
	re := regexp.MustCompile(command)
	return re.ReplaceAllString(s, replacement)
}
