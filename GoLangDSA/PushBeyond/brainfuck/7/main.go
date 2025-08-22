package main

import (
    "os"
    "github.com/01-edu/z01"
)

const SIZE = 2048 

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    var tape [SIZE]byte
    var ptr int
    code := args[0]
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


// >: increment the pointer.
// <: decrement the pointer.
// +: increment the pointed byte.
// -: decrement the pointed byte.
// .: print the pointed byte to standard output.
// [: If the byte at the current pointer is 0, skip forward to the command after the matching ].
// ]: If the byte at the current pointer is not 0, jump back to the command after the matching [.
// // Any other character is treated as a comment and ignored.
