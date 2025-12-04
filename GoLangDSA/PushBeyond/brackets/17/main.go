// $ go run . '(johndoe)' | cat -e
// OK$
// $ go run . '([)]' | cat -e
// Error$
// $ go run . '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
// OK$
// OK$

package main

import (
    "os"
    "fmt"
)


func isBracketBalanced(s string) bool {
    if s == "" {return true}

    var stack []rune
    for _, char := range s {
        if char == '[' || char == '{' || char == '(' {
            stack = append(stack, char)
            continue
        }

        if char == ']' {
            if len(stack) == 0 || stack[len(stack)-1] != '[' {return false}
            stack = stack[:len(stack)-1] // pop
        } else if char == ')' {
            if len(stack) == 0 || stack[len(stack)-1] != '(' {return false}
            stack = stack[:len(stack)-1] // pop
        } else if char == '}' {
            if len(stack) == 0 || stack[len(stack)-1] != '{' {return false}
            stack = stack[:len(stack)-1] // pop
        } 
    }

    return len(stack) == 0
}



func main() {
    args := os.Args[1:]
    if len(args) == 0 {fmt.Println(); return}

    for _, arg := range args {
        if isBracketBalanced(arg) {
            fmt.Println("OK")
        } else {
            fmt.Println("Error")
        }
    }

}


