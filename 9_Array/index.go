package main

import "fmt"

func main() {
	var array = [3]int{1, 2, 3}
	var array2 = [4]int{4, 5, 6}

	var numbers [5]int
	fmt.Println(array)
	fmt.Println(array2)
	fmt.Println(numbers)

	numbers[0] = 10
	fmt.Println(numbers)

	var lovers [2]string
	lovers[0] = "ME"
	lovers[1] = "Only ME"
	fmt.Println(lovers)

	var array3 = [4]string{"Saleh", "Abdulla", "Saleh", "Abdulla"}
	fmt.Println(len(array3))
	fmt.Println(cap(array3))

	for item := 0; item < len(array3); item++ {
		fmt.Println(array3[item])
	}

	for item := range len(array3) {
		fmt.Println(array3[item])
	}

	for item := range len(array3) {
		fmt.Println(array2[item])
	}

	

}
