package main

import "fmt"


func main(){
	fmt.Println(HashCode("14th Avril"))
}

func HashCode(str string) string{
	var result []rune
	for _, char := range str {
		charIntoAscii := (int(char) + len(str)) % 127
		if charIntoAscii < 30 {
			charIntoAscii += 33
		}
		result = append(result, rune(charIntoAscii))
	}

	return string(result)
}













// func HashCode(s string) string {
// 	var end []rune
// 	for _, r := range s {
// 		var ch = (int(r) + len(s)) % 127
// 		if (ch < 30) {ch += 33}
// 		end = append(end, rune(ch))
// 	}
// 	return string(end)
// }
