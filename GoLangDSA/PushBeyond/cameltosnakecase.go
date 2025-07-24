package piscine

//Instructions
// Write a function that converts a string from camelCase to snake_case.

// If the string is camelCase, return the snake_case version of the string.
// For this exercise you need to know that camelCase has two different writing alternatives that will be accepted:

// lowerCamelCase
// UpperCamelCase
// Rules for writing in camelCase:

func CamelToSnakeCase(s string) string {
	// If the string is empty, return an empty string.
	if s == "" {
		return ""
	}

	// if all lowercase return s
	weHaveUpperCase := false
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			weHaveUpperCase = true
			break
		}
	}

	if !weHaveUpperCase {
		return s
	}

	// If the string is not camelCase, return the string unchanged. camaCase => salehAhmedAbdulla
	for i, char := range s {
		if (char >= 'A' && char <= 'Z') && i != 0 {
			// once I encountered an Uppercase, I will check the previous char, if its small, if not return s
			if !(s[i-1] >= 'a' && s[i-1] <= 'z') { //
				return s
			}
		}
	}

	// last char ends with uppercase  // The word does not end on a capitalized letter (CamelCasE).
	if s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z' {
		return ""
	}

	/// Any number or punctuation // Numbers or punctuation are not allowed in the word anywhere (camelCase1).
	for _, char := range s {
		if char >= ' ' && char <= '@' || char >= '[' && char <= '`' || char >= '{' && char <= '~' {
			return ""
		}
	}

	//// No two capitalized letters shall follow directly each other (CamelCAse). Back and after Checker
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			if i == 0 && len(s) > 1 {
				if s[i+1] >= 'A' && s[i+1] <= 'Z' {
					return "" // two Upper Violation
				}
			}
			//----

			if i > 0 && i+1 < len(s) {
				if s[i+1] >= 'A' && s[i+1] <= 'Z' || s[i-1] >= 'A' && s[i-1] <= 'Z' {
					return "" // two Upper Violation
				}
			}
		}
	}

	/// ---- Algo is here
	result := ""
	for i, char := range s {
		if char >= 'A' && char <= 'Z' && i != 0 {
			result += "_"
		}
		result += string(char)
	}

	return result


}
