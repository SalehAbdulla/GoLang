package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	job    string
	salary float64
}

type Programmer struct {
	Person
	isProgrammer bool
}

func (p Person) greeting() {
	fmt.Println("Hello ", p.name)
}

func main() {
	var userOne Person

	userOne.name = "saleh"
	userOne.age = 26
	userOne.job = "Software Engineer"
	userOne.salary = 1000000000

	fmt.Println(userOne)

	userTwo := Person{
		name:   "Saleh",
		age:    26,
		job:    "Software Engineer",
		salary: 10000000000,
	}

	fmt.Println(userTwo)

	saleh := Programmer{
		Person:       userOne,
		isProgrammer: true,
	}

	fmt.Println(saleh)
	saleh.greeting();


	// Anonymous

	me := struct {
		firstName string
		lastName string
	} {
		firstName: "Saleh",
		lastName: "Abdulla",
	}

	fmt.Println(me.firstName);
	fmt.Println(me.lastName);

}
