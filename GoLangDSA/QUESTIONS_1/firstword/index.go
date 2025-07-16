package main

import (
    "fmt"
   
)

func main() {
    fmt.Print(FirstWord("hello there"))
    fmt.Print(FirstWord(""))
    fmt.Print(FirstWord("hello   .........  bye"))
}

// Write a function that takes a string and return a string containing its first word, followed by a newline ('\n').
// A word is a sequence of characters delimited by spaces or by the start/end of the argument.

func FirstWord(s string) string {
    // ...
	if s == "" {
        return s + "\n"
    }

    result := ""
    for _, char := range s {
        if char == ' ' {
            return result + "\n"
        } else {
            result = result + string(char)
        }
    }

    return result + "\n"
}
