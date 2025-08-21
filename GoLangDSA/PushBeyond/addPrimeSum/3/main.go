package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main(){
    args := os.Args[1:]
    if len(args) != 1 {z01.PrintRune('0');z01.PrintRune('\n'); return}
    input := args[0]
    inputToInt, ok := Atoi(input)
    if !ok {z01.PrintRune('0');z01.PrintRune('\n'); return}
    var total int
    for i := inputToInt; i >= 2; i-- {
        if IsPrime(i) {
            total += i
        }
    }
    for _, char := range Itoa(total){
        z01.PrintRune(char)
    }
    z01.PrintRune('\n')
}



func IsPrime(num int) bool {
    
    if num < 2 {return false}

    i := 2
    for i*i <= num {
        if num %i == 0 {return false}
        i++
    }

    return true
}


func Atoi(num string) (int, bool) {
    if num == "0" {return 0, true}
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

    if result < 2 {return -1, false}

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

