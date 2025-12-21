package utils

import (
	"testing"
)

func TestRemoveComment_test(t *testing.T) {
	lines := []string{
		"0",
		"##comment",
		"##start",
		"1 3 5",
		"1 1 0",
		"1 1 8",
		"3 1 5 ",
		"##end",
	}

	expected := []string{
		"0",
		"##start",
		"1 3 5",
		"1 1 0",
		"1 1 8",
		"3 1 5 ",
		"##end",
	}

	filterTheComment := RemoveComments(lines)

	if len(filterTheComment) == len(lines) {
		t.Error("Result of RemoveComments cannot be the same length!")
	}

	for i := 0; i < len(lines); i++ {
		var val1 string
		var val2 string
		if i < len(expected) {val1 = expected[i]}
		if i < len(filterTheComment) {val2 = lines[i]}
		if val1 != val2 {t.Error("Got: %s Expected %s\n", val2, val1)}
	}



}