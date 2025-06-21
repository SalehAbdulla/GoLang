package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("");
	}

	for i := 0; i < 20; i++ {
		fmt.Println("------OUTER-------")
		for j := 0; j < 10; j++ {
			fmt.Println("------INNER------")
		}
	}


}