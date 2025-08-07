package main

import ("github.com/01-edu/z01";"os")

func IsVowel(r rune) bool {
    return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func Print(s string) {
    for _, char := range s {
        z01.PrintRune(char)
    }
}

func main(){
    args := os.Args[1:]
    if len(args) != 1 {return}
    input := args[0]
    if input == "" {return}
    vowelIndex := -1

    for i, char := range input {
        if IsVowel(char) {
            vowelIndex = i
            break
        }
    }

    if vowelIndex == -1 {Print("No vowels\n"); return}

    if vowelIndex == 0 {
        Print(input)
        Print("ay\n")
    } else {
        Print(input[vowelIndex:])
        Print(input[:vowelIndex])
        Print("ay\n")
    }

}