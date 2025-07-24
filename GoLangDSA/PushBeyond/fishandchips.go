package piscine

// Write a function called FishAndChips() that takes an int and returns a string.

// If the number is negative return error: number is negative.
// If the number is non divisible by 2 or 3 return error: non divisible.
// If the number is divisible by 2 and 3, print fish and chips.
// If the number is divisible by 3, print chips.
// If the number is divisible by 2, print fish.

func FishAndChips(n int) string {
	if n < 0 {
		return "number is negative"
	} else if n%3 != 0 || n%2 != 0 {
		return "non divisible"
	} else if n%3 == 0 && n%2 == 0 {
		return "fish and chips."
	} else if n%3 == 0 {
		return "chips"
	} else if n%2 == 0 {
		return "fish"
	}
	return ""
}
