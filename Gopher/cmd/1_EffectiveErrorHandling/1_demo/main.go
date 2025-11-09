package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	TruckIsBusy         = errors.New("truck is currently being used by other department")
	ErrorNotImplemented = errors.New("error is not implemented")
	ErrorTruckNotFound  = errors.New("error truck is not found")
)

// abstract struct
type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo -= 1
	e.battery -= 2
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo += 1
	e.battery += 1
	return nil
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo -= 1
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("processing truck: %+v\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo")
	}
	return nil
}

func main() {

	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "1"}

	// When we working with marsheling JSON
	person := make(map[string]interface{}, 0) // or any 1.18 GO has intoduced
	person["name"] = "Tiago"
	person["age"] = 42

	age, exists := person["age"].(int)
	if !exists {
		log.Fatal("age does not exist")
	}
	log.Println(age)

	if err := processTruck(nt); err != nil {
		log.Fatal(err.Error())
	}

	if err := processTruck(et); err != nil {
		log.Fatal(err.Error())
	}

	log.Println(nt.cargo)
	log.Println(et.battery)

}
