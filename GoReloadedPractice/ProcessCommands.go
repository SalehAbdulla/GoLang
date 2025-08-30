package reloaded

import "log"

func ProcessCommands(input string) string {
	ref := input
	currentCommand := ""
	inCommand := false
	commandStartsAt := 0

	for i := 0; i < len(ref); i++ {
		char := ref[i]

		if char == '(' {
			inCommand = true
			currentCommand = "("
			commandStartsAt = i
			continue
		}

		if inCommand {
			currentCommand += string(char)
			if char == ')' {
				cmd, count := FetchCommand(currentCommand)
				ref = ApplyCommand(ref, cmd, count, commandStartsAt)
				i = -1
				inCommand = false
				currentCommand = ""
			}
		}
	}

	if HasPunctuation(ref) {
		var err error
		ref, err = ReloadPuctuation(ref)
		if err != nil {log.Fatal(err.Error())}
	}

	ref = ReloadVowels(ref)

	return ref
}
