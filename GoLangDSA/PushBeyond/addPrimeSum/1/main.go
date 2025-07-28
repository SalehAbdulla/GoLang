package main

import (
    "os"
    "github.com/01-edu/z01"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
    }
    if atoi(args[0]) < 0 {
        z01.PrintRune('0')
        z01.PrintRune('\n')
    }
    sum := 0
    for i := 2; i <= atoi(args[0]); i++ {
        if IsPrime(i) {
            sum += i
        }
    }
    for _, char := range Itoa(sum) {
        z01.PrintRune(char)
    }
    z01.PrintRune('\n')

}

func atoi(num string) int {
    if num == "0" {return 0}
    isNegative := false
    if num[0] == '-' {
        isNegative = true
        num = num[1:]
    }
    // ------------------
    result := 0

    for _, char := range num {
        result *= 10
        result += int(char - '0')
    }

    // ------------------
    if isNegative{
        return result *-1
    } else {
        return result
    }


}



func Itoa(num int) string {
    if num == 0 {return "0"}

    isNegative := false
    if num < 0 {
        isNegative = true
        num = -num
    }
    result := ""

    /// ---------------

    for num > 0 {
        digit := num % 10
        result = string(rune(digit + '0')) + result
        num = num / 10
    }

    /// ---------------
    if isNegative {
        return "-" + result
    } else {
        return result
    }
}


func IsPrime(num int) bool {
    if num < 2 {return false}
    for i := 2; i*i <= num; i++ {
        if num %i == 0 {
            return false
        }
    }
    return true
}
