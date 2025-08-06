package main
import ("os"; "fmt")

// $ go run . "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
// 1: heavy
// 2: steady
// 3: heavy
// -------------------------

/// ----- Helper function

func contains(s, toFind string) bool {
    for i := 0; i <= len(s)-len(toFind);i++ {
        if toFind == s[i:i+len(toFind)] {return true}
    }
    return false
}

func split(s, splitter string) []string {
    words := []string{}
    word := ""
    for _, char := range s {
        if string(char) != splitter {
            word += string(char)
        } else if word != "" {
            words = append(words, word)
            word = ""
        }
    }    
    if word != "" {words = append(words, word)}
    return words
}

func endsWithRune(s string) bool {
    return (s[len(s)-1] >= 'a' && s[len(s)-1] <= 'z') || (s[len(s)-1] >= 'A' && s[len(s)-1] <= 'Z')
}

func search(exps, texts []string) (result []string) {
    for _, text := range texts {
        for _, exp := range exps {
            if contains(text, exp) {
                if endsWithRune(text) {
                    result = append(result, text)
                } else {
                    result = append(result, text[:len(text)-1])
                }
            }
        }
    }
    return
}



func main(){
    args := os.Args[1:]
    if len(args) != 2 {return}
    regularExp := args[0]
    text := args[1]
    if regularExp == "" || text == "" {return}
    if regularExp[0] != '(' && regularExp[len(regularExp)-1] != ')' {return}
    
    regularExp = regularExp[1:len(regularExp)-1] // get rid of brackets (a) -> a || (a|b) -> a|b


    // Exp to Slice
    var expToSlice []string
    if contains(regularExp, "|") {
        expToSlice = split(regularExp, "|")
    } else {
        expToSlice = append(expToSlice, regularExp)
    }

    // text to Slice
    var textToSlice []string
    textToSlice = append(textToSlice, split(text, " ")...)

    // Search Algorithem
    solution := search(expToSlice, textToSlice)
    for i, solu := range solution {
        fmt.Printf("%d: %s\n", i+1, solu)
    }

}