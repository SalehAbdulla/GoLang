package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func OpenConnection(done chan bool) {
	fmt.Println("Attempting Connection...")

	if rand.Intn(100) > 5 {
		fmt.Println("OOPS! Hanging Connection")
	} else {
		time.Sleep(2 * time.Second)
		fmt.Println("Connnectoin Established!")
	}
}

func openConnectionWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan bool)
	go OpenConnection(done)

	select {
	case <- done:
		fmt.Println("Connection Successful")
	case <- ctx.Done():
		fmt.Println("Connection Timeout!")
	}

}

func main(){
	openConnectionWithTimeout()	
}