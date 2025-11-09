package main

import (
	"testing"
)

func TestTruckOperations(t *testing.T) { // Renamed TestMain to TestTruckOperations for clarity
	t.Run("processTruck methods", func(t *testing.T) {
		t.Run("Should correctly load and unload truck cargo", func(t *testing.T) {
			// Normal Truck initialization: cargo starts at 42
			nt := &NormalTruck{id: "1", cargo: 42} 
			
			// Electric Truck initialization: battery starts at 0.0
			et := &ElectricTruck{id: "1"} 

			// --- Normal Truck Operations ---
			// nt.LoadCargo(): 42 + 1 = 43
			nt.LoadCargo()
			// nt.UnloadCargo(): 43 - 1 = 42
			nt.UnloadCargo()

			// --- Electric Truck Operations ---
			// et.LoadCargo(): battery: 0.0 - 2.0 = -2.0
			et.LoadCargo()
			// et.UnloadCargo(): battery: -2.0 + 1.0 = -1.0
			et.UnloadCargo()

			// ASSERTION 1: Normal Truck
			// The original cargo of 42 is restored after Load (+1) and Unload (-1).
			const expectedNormalCargo = 42 
			if nt.cargo != expectedNormalCargo {
				t.Fatalf("Normal truck cargo mismatch. Expected %d, got %d", expectedNormalCargo, nt.cargo)
			}

			// ASSERTION 2: Electric Truck
			// Battery: (Initial 0) - (Load 2) + (Unload 1) = -1.0
			const expectedElectricBattery = -1.0
			if et.battery != expectedElectricBattery {
				t.Fatalf("Electric truck battery mismatch. Expected %f, got %f", expectedElectricBattery, et.battery)
			}

		})
	})
}