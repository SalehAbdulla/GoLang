package piscine

import (
	"strconv"
)

func FromTo(from int, to int) string {
	if from > 99 || from < 0 || to > 99 || to < 0 {
		return "Invalid\n"
	}
	var result string
	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i != to {
				result += ", "
			}
		}
	} else if from >= to {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0"
			}
			result += strconv.Itoa(i)
			if i != to {
				result += ", "
			}
		}
	}

	return result + "\n"
}
