package reloaded

import "strings"

func RemoveCommand(text string) string {
	from := strings.Index(text, "(")
	to := strings.Index(text, ")")

	if from != -1 && to != -1 && to > from {
		return strings.TrimSpace(text[:from] + text[to+1:])
	}
	return text
}
