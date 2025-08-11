package main

import (
    "os"
    "fmt"
)

type RumanNumeral struct {
    val int
    symbol string
    formula string
}

var numerals = []RumanNumeral {
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
        for num >= numeral.val {
            symbol += numeral.symbol
            if len(formula) != 0 {formula += "+"}
            formula += numeral.formula
            num -= numeral.val
        }
    }

    return symbol, formula
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}

    input := args[0]
    inputToInt, ok := stringToInt(input)
    if !ok {fmt.Printf("ERROR: cannot convert to roman digit\n"); return}

    getRumanSymbol, getRumanFormula := GetRuman(inputToInt)
    fmt.Println(getRumanFormula)
    fmt.Println(getRumanSymbol)

}

func stringToInt(s string) (int, bool) {
    if s == "0" || s == "" {return -1, false}

    var result int
    for _, char := range s {
        if char < '0' || char > '9' {
            return -1, false
        } else {
            result *= 10
            result += int(char - '0')
        }
    }
    if result < 0 || result > 3999 {return -1, false}
    return result, true
}
