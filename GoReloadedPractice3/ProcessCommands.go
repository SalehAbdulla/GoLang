package reloaded

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
				cmd, count := ParseCommand(currentCommand)
				ref = ApplyCommands(ref, cmd, count, commandStartsAt)
				i = -1
				inCommand = false
				currentCommand = ""
			}
		}
	}

	if HasSymbol(ref) {
		ref  = FixPunctuation(ref)
	}

	ref = FixVowels(ref)

	return ref
}
