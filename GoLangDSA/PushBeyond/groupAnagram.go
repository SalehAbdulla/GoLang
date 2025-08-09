package piscine

import "fmt"

func GroupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]byte][]string)
	for _, str := range strs {
		hashMap[GetFrequency(str)] = append(hashMap[GetFrequency(str)], str)
	}
	var result [][]string
	for _, value := range hashMap{
		result = append(result, value)
	}
	fmt.Println(result)
	return [][]string{}
}

func GetFrequency(s string) [26]byte {
	var result [26]byte
	for _, c := range s {
		result[c-'a']++
	}
	return result
}
