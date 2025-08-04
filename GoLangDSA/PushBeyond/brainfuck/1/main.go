package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){    
    args := os.Args[1:]
    if len(args) != 1 {return}
    
    code := args[0]
    tape := [2048]byte{}
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
                if tape[ptr] == 0 { // No loop when current value is zero
                    loop := 1
                    for loop > 0 {
                        i++ // Jump Forward
                        if code[i] == '[' {loop++} // Another nested loop, skip
                        if code[i] == ']' {loop--} // we found the exist bracket
                    }
                }
            case ']':
                if tape[ptr] != 0 { // no Exist if current cell is not equal to zero
                    loop := 1
                    for loop > 0{
                        i-- // Go Backward
                        if code[i] == ']'{ // another exit loop Bracket, skip
                            loop++
                        }
                        if code[i] == '[' {
                            loop--
                        }
                    }
                }
        
        }
    }
}
