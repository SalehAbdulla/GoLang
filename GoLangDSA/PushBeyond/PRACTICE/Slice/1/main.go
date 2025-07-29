package main

import (
    "fmt"
    "piscine"
)

func main(){
    a := []string{"coding", "algorithm", "ascii", "package", "golang"}
    fmt.Printf("%#v\n", piscine.Slice(a, 1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
    fmt.Printf("%#v\n", piscine.Slice(a, -3))
    fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
    fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
}

func Slice(a []string, nbrs... int) []string {
	if len(a) == 0 || len(nbrs) == 0 {return nil}

	l := len(a)

	starts := nbrs[0]
	if starts < 0 {starts += l}

	ends := l // by default till the end
	
	if len(nbrs) > 1 {
		ends = nbrs[1]
		if ends < 0 {ends += l}
	}

	/// ---- Validation
	if starts < 0 {starts = 0}
	if ends < 0 {ends = 0}
	if starts > l {starts = l}
	if ends > l {ends = l}
	if starts > ends {return nil}
	/// ---- Validation

	return a[starts:ends]
}