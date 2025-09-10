package reloaded

import (
	"log"
	"regexp"
	"strings"
	"unicode"
)

func FixPunctuation(ref string) string {
	if strings.TrimSpace(ref) == "" {log.Fatal("input is empty or whitespace")}

	for _, char := range ref {
		if !unicode.IsPrint(char) && !unicode.IsSpace(char) {
			log.Fatal("input has non-printable character")
		}
	}

	// 3 pattrens -> 
	// 1) remove spaces before puctuation
	remSpacesBefPunc := `\s+([:;.,?!])`
	// 2) only one space after punctuation
	remSpacesAftPunc := `([,.:;?!])([^\s;:.,?!])`
	// 3) collapse spaces
	collapseSpaces := `\s+`

	ref = compile(ref, remSpacesBefPunc, `$1`)
	ref = compile(ref, remSpacesAftPunc, `$1 $2`)
	ref = compile(ref, collapseSpaces, ` `)

	return ref
}


func compile(ref, pattren, repl string) string {
	re, err := regexp.Compile(pattren)
	if err != nil {log.Fatal(err.Error())}
	return re.ReplaceAllString(ref, repl)
}