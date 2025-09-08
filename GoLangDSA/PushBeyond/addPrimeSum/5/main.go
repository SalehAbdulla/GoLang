package main

import (
    "github.com/01-edu/z01"
    "os"
)

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func Atoi(num string) (int, bool) {
    if num == "" {return -1, false}
    if num == "0" {return -1, false}

    isNegative := false
    if num[0] == '-' {
        isNegative = true
        num = num[1:]
    }

    var result int
    for _, char := range num {
        if char < '0' || char > '9' {
            return -1, false
        } else {
            result *= 10
            result += int(char - '0')
        }
    }

    if isNegative {
        return result *-1, true
    }
    return result, true
}

func Itoa(num int) string {
    if num == 0 {return "0"}
    isNegative := false
    if num < 0 {
        isNegative = true
        num = -num
    }

    var result string

    for num > 0 {
        digit := num %10
        result = string(rune(digit + '0')) + result
        num /= 10
    }

    if isNegative {
        return "-" + result
    }
    return result

}


func main(){
    args := os.Args[1:]
    if len(args) != 1 {Print("0\n"); return}
    num := args[0]
    numToInt, ok := Atoi(num)
    if !ok {Print("0\n"); return}
    var total int
    for i := numToInt; i >= 2; i-- {
        if IsPrime(i) {
            total += i
        }
    }
    totalToString := Itoa(total)
    Print(totalToString)
    Print("\n")
}


func IsPrime(num int) bool {
    i := 2
    for i*i <= num {
        if num %i == 0 {
            return false
        }
        i++
    }
    return true
}