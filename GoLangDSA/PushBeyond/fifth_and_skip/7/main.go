package piscine

//	fmt.Print(piscine.FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
//  abcde ghijk mnopq stuwx z$

// If there are spaces in the middle of a word, 
// ignore them and get the first character after the spaces until you reach a length of 5.


func FifthAndSkip(str string) string {
    if str == "" {return "\n"}
    if len(str) < 5 {return "Invalid Input\n"}
    var clearStr []rune
    for _, char := range str {
        if char != ' ' {
            clearStr = append(clearStr, char)
        }
    }
    var result []rune
    counter := 0
    for _, char := range clearStr {
        if counter == 5 {
            counter = 0
            result = append(result, ' ')
            continue
        }
        counter++
        result = append(result, char)
    }

    return string(result) + "\n"
}