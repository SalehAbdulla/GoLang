package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync" // Required for concurrent processing
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
	// Note: %v in Printf for pointers shows the address, %+v shows struct fields.
	fmt.Printf("Processing truck ID: %s\n", truck) // We can't access ID here directly without type assertion

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo: %w", err)
	}
	return nil
}

// CORRECTED: The processFleet function handles its own WaitGroup internally.
func processFleet(trucks []Truck) error {
	var wg sync.WaitGroup
	
	for _, t := range trucks {
		wg.Add(1)

		// Launch a new goroutine for concurrent processing
		go func(t Truck){
			defer wg.Done()
			// Errors from processTruck are logged but not returned from processFleet 
			// because the goroutine cannot return a value back through the main flow easily.
			if err := processTruck(t); err != nil {
				log.Printf("Error processing truck: %v", err)
			}
		}(t)
	}

	// CORRECTED: wg.Wait() (capital W) blocks until all goroutines finish.
	wg.Wait() 

	return nil
}

func main() {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "userID", 42)

	// REMOVED: The unnecessary sync.WaitGroup from main.
	fleet := []Truck{
		&NormalTruck{id: "NT1", cargo: 0},
		&ElectricTruck{id: "ET1", cargo: 0, battery: 100},
		&NormalTruck{id: "NT2", cargo: 0},
		&ElectricTruck{id: "ET2", cargo: 0, battery: 100}, 
	}

	// CORRECTED: Call processFleet with only the fleet slice.
	if err := processFleet(fleet); err != nil {
		log.Printf("Error processing fleet: %v\n", err)
	}

	// Display final state of the fleet after concurrent processing
	fmt.Println("\n--- Final Fleet State ---")
	for _, t := range fleet {
		fmt.Printf("%+v\n", t)
	}

	fmt.Println("All Trucks processed successfully!")
}