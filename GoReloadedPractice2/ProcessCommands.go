package reloaded

import (
	"log"
)

// This function is gonna capture () pairs
func ProcessCommands(s string) string {
	ref := s
	currentCommand := ""
	cmdIndex := 0
	inCommand := false

	for i := 0; i < len(ref); i++ {
		if ref[i] == '(' {
			inCommand = true
			currentCommand = "("
			cmdIndex = i
		}
		if inCommand {
			currentCommand += string(ref[i])
			if ref[i] == ')' {
				cmd, count := ParseCommand(currentCommand)
				ref = ApplyCommand(ref, cmd, count, cmdIndex)
				inCommand = false
				cmdIndex = 0
				currentCommand = ""
				i = -1 // reset loop
			}
		}
	}

	if HasSymbol(ref) {
		var err error
		ref, err = FixPunctuation(ref)
		if err != nil {log.Fatal(err.Error())}
	}

	ref = FixVowels(ref)

	return ref
}