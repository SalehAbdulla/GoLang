package main

import (
    "os"
    "fmt"
)

type NumeralType struct {
    Val int
    Symbol string
    Formula string
}

// I value X-ray, Let's count Doctor Money
var numerals = []NumeralType {
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

func ConvertToRuman(n int) (string, string) {
    symbol := ""
    formula := ""
    for _, numeral := range numerals {
        for n >= numeral.Val {
            symbol += numeral.Symbol
            if formula != "" {formula += "+"}
            formula += numeral.Formula
            n = n - numeral.Val 
        }
    }
    return symbol, formula
}

func main(){
    args := os.Args[1:]
    if len(args) == 0 {fmt.Println("ERROR: cannot convert to roman digit"); return}
    num, ok := StringToInt(args[0])
    if num == -1 || !ok {fmt.Println("ERROR: cannot convert to roman digit"); return}
    if num <= 0 || num > 3999 {fmt.Println("ERROR: cannot convert to roman digit"); return}
    // ----------
    numToRumanSymbol, numToRumanFormula := ConvertToRuman(num)
    fmt.Println(numToRumanFormula)
    fmt.Println(numToRumanSymbol)
}

func StringToInt(num string) (int, bool) {
    if num == "" {return -1, false}
    if num[0] == '-' {return -1, false}
    result := 0
    for _, char := range num {
        if char >= '0' && char <= '9' {
            result *= 10
            result = int(char - '0') + result
        } else {
            return -1, false
        }
    }
    if result > 3999 || result < 0 {return -1, false}
    return result, true
}