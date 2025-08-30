package reloaded

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func ReloadPuctuation(s string) (string, error) {
	if strings.TrimSpace(s) == "" {
		return s, fmt.Errorf("input string is empty or whitespace")
	}

	for _, r := range s {
		if !unicode.IsPrint(r) && !unicode.IsSpace(r) {
			return s, fmt.Errorf("input contains non-printable characters")
		}
	}
	
	removeSpacesBeforePunc := `\s+([,.!?;:])`
	addOneSpaceAfterPunc := `([,.!?;:])([^\s,.!?;:])`
	collapseSpacesToOne := `\s+`
	
	s = replace(s, removeSpacesBeforePunc, `$1`)
	s = replace(s, addOneSpaceAfterPunc, `$1 $2`)
	s = replace(s, collapseSpacesToOne, " ")
	s = strings.TrimSpace(s)

	return s, nil
}

func replace(s, pattern, replacement string) string {
	re, err := regexp.Compile(pattern)
	if err != nil {log.Fatal("Constant Pattern Failed (panic programmer mistake)")}
	return re.ReplaceAllString(s, replacement)
}
