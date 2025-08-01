package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
)

func ClearEmpty(s []string) (result []string) {
    for _, val := range s {
        if val != "" {result = append(result, val)}
    }
    return
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {fmt.Println("Error"); return }
    if len(args[0]) == 1 {fmt.Println(args[0]); return}
    
    operand := strings.Split(args[0], " ")
    operand = ClearEmpty(operand)

    stack := []int{}

    for _, char := range operand {
        charToInt, err := strconv.Atoi(char)
        // This means it a valid number, then I going to add it to stack
        if err == nil {
            stack = append(stack, charToInt)
            continue
        }
        // If it is not a digit, it must be operation. 
        // First we check if we have two existed digits for operation
        length := len(stack)
        if length < 2 {fmt.Println("Error"); return}

        switch char {
            case "+":
                // operation made on last two digits, update the prevLast
                stack[length - 2] += stack[length - 1]
                stack = stack[:length -1] // get rid of last element
            case "-":
                // operation made on last two digits, update the prevLast
                stack[length - 2] -= stack[length - 1]
                stack = stack[:length -1] // get rid of last element
            case "*":
                // operation made on last two digits, update the prevLast
                stack[length - 2] *= stack[length - 1]
                stack = stack[:length -1] // get rid of last element
            case "/":
                // operation made on last two digits, update the prevLast
                stack[length - 2] /= stack[length - 1]
                stack = stack[:length -1] // get rid of last element
            case "%":
                // operation made on last two digits, update the prevLast
                stack[length - 2] %= stack[length - 1]
                stack = stack[:length -1] // get rid of last element
            default:
                fmt.Println("Error") // if none of these, then this is invalid char
                return
        }
        
    }
    
    if len(stack) == 1{
        fmt.Println(stack[0])
    } else {
        fmt.Println("Error")
    }
    
}