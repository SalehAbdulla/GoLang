package main

import (
    "os"
    "github.com/01-edu/z01"
)

// >: increment the pointer.
// <: decrement the pointer.
// +: increment the pointed byte.
// -: decrement the pointed byte.
// .: print the pointed byte to standard output.
// [: If the byte at the current pointer is 0, skip forward to the command after the matching ].
// ]: If the byte at the current pointer is not 0, jump back to the command after the matching [.
// Any other character is treated as a comment and ignored.

const BYTESIZE uint64 = 2048

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    code := args[0]
    tape := [BYTESIZE]byte{}
    ptr := 0

    for i := 0; i < len(code); i++ {
        switch code[i]{
            case '>':
                ptr++
            case '<':
                ptr--
            case '-':
                tape[ptr]--
            case '+':
                tape[ptr]++
            case '.':
                z01.PrintRune(rune(tape[ptr]))
            case '[': // If the byte at the current pointer is 0, skip forward
                if tape[ptr] == 0 {
                    loop := 1
                    for loop > 0 {
                        i++
                        if code[i] == '[' {loop++} // Edge case, nested loop
                        if code[i] == ']' {loop--} // exit loop
                    }
                }
            case ']': // If the byte at the current pointer is not 0, jump back to [
                if tape[ptr] != 0 {
                    loop := 1
                    for loop > 0{
                        i--
                        if code[i] == ']' { // nested closing loop
                            loop++
                        }
                        if code[i] == '[' {loop--} // We reached [
                    }
                }
        }
    }
}
