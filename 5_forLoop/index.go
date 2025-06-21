package main

import (
	"fmt"
)

func main () {

	for i := range 5 {
			fmt.Println(i);
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i);
	}

	for j := 0; j <= 10; j++ {

		if j == 5 {
			continue;
		}

		fmt.Println(j);
	}


 
}