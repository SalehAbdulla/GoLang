package main

import (
	"fmt"
)


func main(){
	strs := []string {"act","pots","tops","cat","stop","hat"}
	fmt.Println(groupAnagrams(strs));
}

func groupAnagrams(strs []string) [][]string {
    res := make(map[[26]int][]string)

    for _, s := range strs {
        var count [26]int
        for _, c := range s {
            count[c-'a']++
        }
        res[count] = append(res[count], s)
    }

    var result [][]string
    for _, group := range res {
        result = append(result, group)
    }
    return result

}