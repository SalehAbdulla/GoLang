package piscine

func FifthAndSkip(str string) string {
    if str == "" {return "\n"}
    var clearStr string
    for _, char := range str {if char != ' ' {clearStr += string(char)}}
    if len(clearStr) < 5 {return "Invalid Input\n"}
    // abcdefghijklmnopqrstuwxyz
    // abcde ghijk mnopq stuwx z$
    var buffer string
    counter := 0
    for _, char := range clearStr {
        if counter == 5 {
            counter = 0
            buffer += " "
            continue
        }
        buffer += string(char)
        counter++
    }
    return buffer + "\n"
}
