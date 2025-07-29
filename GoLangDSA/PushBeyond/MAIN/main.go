package main

import (
	"fmt"
	"strings"
)

func main() {
	arrStr := "[5, 6, 7, 8, 9]"
	numStrs := strings.Split(strings.Trim(arrStr, "[]"), ",")
	fmt.Println(numStrs)

}
