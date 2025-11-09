package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrorNotImplemented = errors.New("not implemented")
	TrackNotFound       = errors.New("track not found")
)

type Truck struct {
	id    string
	cargo int
}

func (t *Truck) LoadCargo() error {
	return nil
}

func (t *Truck) UnloadCargo() error {
	return nil
}

// HandleTruck handles the loading and unloading of a truck
func HandleTruck(truck Truck) error {
	fmt.Println("Processing truck: %s\n", truck.id)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error cargo: %w", err)
	}

	fmt.Printf("Process Truck: %s\n", truck)
	return ErrorNotImplemented
}

func main() {
	trucks := []Truck{
		{id: "Truck-1"},
		{id: "Truck-2"},
		{id: "Truck-3"},
	}

	for _, truck := range trucks {
		if err := HandleTruck(truck); err != nil {
			switch err {
			case ErrorNotImplemented:
				// we do this
				log.Fatal(ErrorNotImplemented)
				return
			case TrackNotFound:
				// We do this
				log.Fatal(TrackNotFound)
				return
			default:
				log.Fatal(err.Error())
			}

		}
	}

}
