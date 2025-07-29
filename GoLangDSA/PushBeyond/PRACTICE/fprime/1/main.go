package main

//- Factors must be displayed in ascending order and separated by `*`.

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) != 1 {return}
	if Atoi(args[0]) < 2 {return}
	// -----------
	num := Atoi(args[0])
	result := []int{}

	i := 2
	for i*i <= num {
		for num%i == 0 {
			result = append(result, i)
			num /= i
		}
		i++
	}

	if num > 1 {
		result = append(result, num) // it must be a prime
	}

	for i, num := range result {
		fmt.Print(num)
		if i != len(result) -1 {
			fmt.Print("*")
		}
	}
	fmt.Println()

}


func Atoi(num string) int {
	if num == "" {return -1}
	if num == "0" {return 0}
	isNegative := false
	if num[0] == '-' {
		isNegative = true
		num = num[1:]
	}
	result := 0
	for _, char := range num {
		if char >= '0' && char <= '9' {
			result *= 10
			result += int(char - '0')
		} else {
			return -1
		}
	}
	if isNegative {
		return result * -1
	}
	return result
}
