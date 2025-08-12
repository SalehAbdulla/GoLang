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

    var tape [TAPE_SIZE]byte
    ptr := 0

    for i := 0; i < len(code); i++ {
        switch code[i]{
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
                for loop > 0{
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