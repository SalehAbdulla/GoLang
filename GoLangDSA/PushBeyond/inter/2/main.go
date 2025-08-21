package main 

import(
    "github.com/01-edu/z01"
    "os"
)

func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    hashMap1 := make(map[rune]bool)
    hashMap2 := make(map[rune]bool)
    printed := make(map[rune]bool)
    for _, char := range args[0]{hashMap1[char]=true}
    for _, char := range args[1]{hashMap2[char]=true}

    for _, char := range args[0] {
        if hashMap2[char] && !printed[char] {z01.PrintRune(char); printed[char]=true}
    }

    for _, char := range args[1] {
        if hashMap1[char] && !printed[char] {z01.PrintRune(char); printed[char]=true}
    }
    z01.PrintRune('\n')
}