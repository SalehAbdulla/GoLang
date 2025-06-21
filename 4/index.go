package main

import (
	"fmt"
	"strings"
)

func main() {
	findChar := "Golang programming language";
	fmt.Println(strings.Contains(findChar, "Go"));

	// Upper Lower

	stringOne := "This is Some Random Text";
	fmt.Println(strings.ToUpper(stringOne));
	fmt.Println(strings.ToLower(stringOne));

	
}