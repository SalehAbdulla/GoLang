package main


import (
	"fmt"
)

func main() {
	userData := map[string]int {
		"Saleh": 26,
		"Abdulla": 24,
		"Ahmed": 29,
	}

	fmt.Println(userData["Saleh"]);

	for key, value := range userData {
		fmt.Println(key, value);
	}


	newStudent := map[int]string{
		1: "Saleh",
		2: "Abdulla",
		3: "Mahdi",
	}

	for key, value := range newStudent {
		fmt.Println(key , " " , value);
	}

	


}