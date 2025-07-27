package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 2 {
		return
	}

	factors := []int{}
	i := 2
	for i*i <= n {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
		i++
	}

	if n > 1 {
		factors = append(factors, n)
	}

	for i, v := range factors {
		fmt.Print(v)
		if i != len(factors)-1 {
			fmt.Print("*")
		}
	}
	
	fmt.Println()
}
