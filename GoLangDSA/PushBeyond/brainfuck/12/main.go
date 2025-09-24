package main

import (
    "os"
    "github.com/01-edu/z01"
)
const TAPE_SIZE = 2048 
func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    code := args[0]
    ptr := 0
    tape := make([]byte, TAPE_SIZE)

    for i := 0; i < len(code); i++ {
        switch code[i] {
            case '+':
                tape[ptr]++
            case '-':
				if tape[ptr] > 0 {
					tape[ptr]--
				}
            case '>':
                ptr++
            case '<':
                ptr--
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
