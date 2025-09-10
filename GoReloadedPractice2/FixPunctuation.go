package reloaded

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func FixPunctuation(ref string) (string, error) {
	if strings.TrimSpace(ref) == "" {return ref, fmt.Errorf("input is empty or whitespace")}
	ref = strings.TrimSpace(ref)

	for _, char := range ref {
		if !unicode.IsPrint(char) && !unicode.IsSpace(char) {
			return ref, fmt.Errorf("input has non-printable character")
		}
	}

	// 3 Pattrens needed, before , && one space after,  && collapse crazy spaces
	removeSpacesAfter := `\s+([:;.,?!])`
	removeSpacesBefore := `([:;.,?!])([^\s.,?!:;])`
	collapseSpaces := `\s+`

	ref = compile(ref, removeSpacesAfter, "$1")
	ref = compile(ref, removeSpacesBefore, "$1 $2")
	ref = compile(ref, collapseSpaces, " ")

	return ref, nil
}


func compile(ref, pattren, repl string) string {
	re, err := regexp.Compile(pattren)
	if err != nil {log.Fatal(err.Error())}
	return re.ReplaceAllString(ref, repl)
}
