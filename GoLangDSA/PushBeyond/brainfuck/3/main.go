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

        switch code[i] {
            case '>':
                ptr = (ptr + 1) % TAPE_SIZE
            case '<':
                ptr = (ptr - 1 + TAPE_SIZE) % TAPE_SIZE
            case '+':
                tape[ptr]++
            case '-':
                tape[ptr]--
            case '.':
                z01.PrintRune(rune(tape[ptr]))
            case '[':
            // [: If the byte at the current pointer is 0, skip forward to the command after the matching ].
                if tape[ptr] == 0 {
                    loop := 1
                    for loop > 0 {
                        i++
                        if code[i] == '[' {loop++} // catch another loop
                        if code[i] == ']' {loop--} // catch ending loop
                    }
                }
            case ']':
                if tape[ptr] != 0 {
                    loop := 1
                    for loop > 0 {
                        i--
                        if code[i] == ']' {loop++} // Another ending loop
                        if code[i] == '[' {loop--} // catch starting loop
                    }
                }
            
        }
    }

}
