package parsing

import "strings"

func RemoveComments(inStr string) string {
	// First things we remove \r
	inStr = strings.ReplaceAll(inStr, "\r", "")

	splitStr := strings.Split(inStr, "\n")
	var noComment string

	for i, line := range splitStr {
		if !strings.HasPrefix(line, "##") && strings.HasPrefix(line, "#") {
			continue
		}

		noComment += line
		if i != len(splitStr)-1 {
			noComment += "\n"
		}
	}

	return noComment
}
