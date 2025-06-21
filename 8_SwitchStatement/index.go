package main

import "fmt"

func main() {

	today := 1;

	switch today {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	default:
		fmt.Println("Not Sunday or Monday")
	}

}
