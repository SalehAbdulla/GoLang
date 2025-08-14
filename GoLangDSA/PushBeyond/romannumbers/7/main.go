package main

import (
    "os"
    "fmt"
)

type RumanNumeral struct {
    Val int
    Symbol string
    Formula string
}

var numerals = []RumanNumeral{
    {1000,"M","M"},
    {900,"CM","(M-C)"},
    {500,"D","D"},
    {400,"CD","(D-C)"},
    {100,"C","C"},
    {90,"XC","(C-X)"},
    {50,"L","L"},
    {40,"XL","(L-X)"},
    {10,"X","X"},
    {9,"IX","(X-I)"},
    {5,"V","V"},
    {4,"IV","(V-I)"},
    {1,"I","I"},
}

func GetRuman(num int) (string, string) {
    var symbol string
    var formula string

    for _, numeral := range numerals {
        for num >= numeral.Val {
            symbol += numeral.Symbol
            if formula != "" {formula += "+"}
            formula += numeral.Formula
            num -= numeral.Val
        }
    }
    return symbol, formula
}

func StringToInt(num string) (int, bool) {
    if len(num) == 0 || num == "" || num == "0" || num[0] == '-' {return -1, false}
    var result int
    for _, char := range num {
        if char < '0' || char > '9' {
            return -1, false
        } else {
            digit := int(char - '0')
            result *= 10
            result += digit
        }
    }
    if result < 1 || result > 3999 {return -1, false}
    return result, true
}



func main() {
    args := os.Args[1:]
    if len(args) == 0 {return}
    inputToInt, ok := StringToInt(args[0])
    if !ok {fmt.Println("ERROR: cannot convert to roman digit"); return}
    getSymbol, getFormula := GetRuman(inputToInt)
    fmt.Println(getFormula)
    fmt.Println(getSymbol)

}




