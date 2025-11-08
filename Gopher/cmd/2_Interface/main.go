package main

import (
	"errors"
	"fmt"
	"log"
)

// Define errors
var (
	ErrorNotImplemented = errors.New("not implemented")
	TruckNotFound       = errors.New("truck not found")
)

// Truck interface defines expected behaviors
type Truck interface {
	ID() string
	LoadCargo() error
	UnloadCargo() error
}

// ElectricTruck implementation
type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) ID() string {
	return e.id
}

func (e *ElectricTruck) LoadCargo() error {
	fmt.Println("Electric truck loading cargo...")
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	fmt.Println("Electric truck unloading cargo...")
	return nil
}

// NormalTruck implementation
type NormalTruck struct {
	id    string
	cargo int
	fuel  float64
}

func (t *NormalTruck) ID() string {
	return t.id
}

func (t *NormalTruck) LoadCargo() error {
	fmt.Println("Normal truck loading cargo...")
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	fmt.Println("Normal truck unloading cargo...")
	return nil
}

// HandleTruck handles the loading and unloading of a truck
func HandleTruck(truck Truck) error {
	fmt.Printf("Processing truck: %s\n", truck.ID())

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("error loading cargo: %w", err)
	}

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	fmt.Printf("Processed Truck: %s successfully\n", truck.ID())
	return ErrorNotImplemented
}

func main() {
	trucks := []Truck{
		&NormalTruck{id: "T-001", cargo: 50, fuel: 80.0},
		&ElectricTruck{id: "E-001", cargo: 20, battery: 90.0},
	}

	for _, truck := range trucks {
		if err := HandleTruck(truck); err != nil {
			switch {
			case errors.Is(err, ErrorNotImplemented):
				log.Println("Feature not implemented yet:", err)
			case errors.Is(err, TruckNotFound):
				log.Println("Truck not found:", err)
			default:
				log.Println("Unexpected error:", err)
			}
		}
	}

}
