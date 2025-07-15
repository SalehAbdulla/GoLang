package main

import (
	"fmt"
)

func main(){
	nums := []int {1,2,4,6}
	productArrayExceptSelf(nums)
}


func productArrayExceptSelf(nums []int) []int {
	// --- Makes an array of ones [1 1 1 1]
	res := make([]int, len(nums))

	for i := range res {
		res[i] = 1
	}

	prefix := 1

	for i := 0; i < len(nums); i++ {
		// {1, 2, 4, 6} 1 * 1 | 1 * 1 | 
		res[i] = prefix // product of all elements before
		prefix *= nums[i] // prepare the product for next iteration
	}
	

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}


	fmt.Println(res)
	return res

}

