package utils

import (
    "reflect"
    "testing"
)

var asciiMapSlice = []string{
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

var expected = map[rune][]string{
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

func TestGetAsciiArt(t *testing.T) {
    got := GetAsciiMap(asciiMapSlice)

    if !reflect.DeepEqual(got, expected) {
        t.Errorf("GetAsciiMap(asciiMapSlice) = %#v; want %#v", got, expected)
    }
}
