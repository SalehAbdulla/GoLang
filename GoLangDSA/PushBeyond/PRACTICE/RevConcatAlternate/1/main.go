package main


func RevConcatAlternate(slice1,slice2 []int) []int {
	r1, r2 := reverse(slice1), reverse(slice2)
	res := []int{}
	
	if len(r1) > len(r2) {
		res = append(res, r1[:len(r1)-len(r2)]...) // take off addition ele, and save them to res
		r1 = r1[len(r1)-len(r2):] // get rid off the element
	} else if len(r2) > len(r1) {
		res = append(res, r2[:len(r2)-len(r1)]...) // take off addition ele, and same them to res
		r2 = r2[len(r2)-len(r1):] // get rid off the element
	}

	for i := range r1 {
		res = append(res, r1[i], r2[i])
	}

	return res
}

func reverse(s []int) []int {
	res := []int{}
	for i := range s {
		res = append(res, s[len(s)-1-i]) // starts from the end, append
	}
	return res
}