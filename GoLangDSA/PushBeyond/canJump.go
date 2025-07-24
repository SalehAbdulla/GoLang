package piscine

func CanJump(input []uint) bool {
	if len(input) == 0 {
		return false
	}
	
	if len(input) == 1 {
		return true
	}

	for pos := 0; pos < len(input); {
		steps := int(input[pos])

		// If we're at the last index
		if pos == len(input) -1 {
			return true
		}

		// If stuck (cannot move forward)
		if steps == 0 {
			return false
		}

		// Move forward exactly 'steps' positions
		pos += steps
	}

	// After the loop, check if we landed exactly on the last index
	return false
}

