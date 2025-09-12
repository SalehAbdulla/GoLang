package piscine

func LastWord(s string) string{
    if s == " " || s == "" {return "\n"}
    i := len(s) - 1
    for s[i] == ' ' {
        i--
    }
    ends := i
    for s[i] != ' '{
        i--
    }
    starts := i
    return s[starts+1:ends + 1] + "\n"
}