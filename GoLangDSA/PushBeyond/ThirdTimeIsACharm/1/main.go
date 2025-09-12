package piscine


func ThirdTimeIsACharm(str string) string {
    var result string
    for i := 2; i < len(str); i+=3{
        result += string(str[i])
    }
    return result + "\n"
}
	