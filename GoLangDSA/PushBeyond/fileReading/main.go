package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1:][0]
	file, error := os.Open(fileName)
	
	if error != nil {
		fmt.Println(error)
	} else {
		data := make([]byte, 20)
		file.Read(data)
		fmt.Println(string(data))
		file.Close()
	}
}
