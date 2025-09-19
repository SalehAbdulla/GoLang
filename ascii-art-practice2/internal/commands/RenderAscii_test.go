package commands

import (
	"reflect"
	"strings"
	"testing"
)

var asciiMap = map[rune][]string{
	32: {
		" _  ",
		"| | ",
		"| | ",
		"| | ",
		"|_| ",
		"(_) ",
		"    ",
		"    ",
	},
}

var expected = []string{
    " _  ",
    "| | ",
    "| | ",
    "| | ",
    "|_| ",
    "(_) ",
    "    ",
    "    ",
    "",
}

func TestRenderAscii(t *testing.T) {
	var builder strings.Builder
	RenderAscii("!", asciiMap, &builder)
	got := builder
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("RenderAscii(`!`, asciiMap, &builder) = %#v;\n want %#v", got, expected)
	}
}
