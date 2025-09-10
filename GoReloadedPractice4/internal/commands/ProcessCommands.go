package commands

func ProcessCommands(input string) string {
	ref := input
	ref = Replace(ref, `(?i)/(+\s*hex|bin|cap|low|up\s*)(\s*,\s*\d+\s*)?\s*)+`)

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
				cmd, count, err := commands.ParseCommand(currentCommand)
				if err == nil {
					ref = applyCommand(ref, cmd, count, commandStartsAt)
				} else {
					inCommand = false
					continue
				}
				i = -1
				inCommand = false
				currentCommand = "" // just for clarity
			}
		}
	}

	


}