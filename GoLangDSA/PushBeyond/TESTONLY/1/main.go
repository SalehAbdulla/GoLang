package main

import "fmt"

func main() {
	i, j := 42, 2701
	fmt.Println(&i, &j)


	p := &i
	fmt.Println(*p) // Get its value
}

// Their address, 0xc00000a0e8 0xc00000a100
// There are two usage of *

// Pointer type with _______ base -- int base for example
// or operator type to get the orginal value pointing to