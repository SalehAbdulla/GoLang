package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50};

	fmt.Println(nums[0:4]);
	nums = append(nums, 60);
	fmt.Println(nums);

	num1 := make([]int, 5, 40);
	fmt.Println(num1);

	num1[0] = 10;
	num1[1] = 10;
	num1[2] = 10;
	num1[3] = 10;
	num1[4] = 10;

	fmt.Println(num1);
	fmt.Println(len(num1));
	fmt.Println(cap(num1));

	//--------------

	people := []string {"Saleh", "Ahmed", "Abdulla"};

	for i:= range len(people) {
		fmt.Println(people[i]);
	}

	
	


}