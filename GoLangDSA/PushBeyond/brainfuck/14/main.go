package main

import (
    "github.com/01-edu/z01"
    "os"
)

var ptr int
var tape [2048]byte

func print(s string) {
    for _, char := range s {
        z01.PrintRune(char);
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}

    code := args[0];

    for i := 0; i < len(code); i++ {
        switch code[i] {
            case '>':
                ptr++
            case '<':
                ptr--
            case '+':
                tape[ptr]++
            case '-':
                tape[ptr]--
            case '.':
                z01.PrintRune(rune(tape[ptr]))
            case '[':
                if tape[ptr] == 0 {
                    loop := 1
                    for loop > 0 {
                        i++
                        if code[i] == '[' {loop++}
                        if code[i] == ']' {loop--}
                    }
                }
            case ']':
                if tape[ptr] != 0 {
                    loop := 1
                    for loop > 0 {
                        i--
                        if code[i] == ']' {loop++}
                        if code[i] == '[' {loop--}
                    }
                }
        }
    }
}
