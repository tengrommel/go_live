package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()
	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
	}
	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.GetVehicle()
	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 4 and they were %d\n", bike.Wheels)
	}
	if bike.Structure != "Motorbike" {
		t.Errorf("Structure on a bike must be 'Motorbike' and was %s\n", bike.Structure)
	}
	if bike.Seats != 2 {
		t.Errorf("Seats on a bike must be 2 and they were %d\n", bike.Seats)
	}

	busBuilder := &BusBuilder{}
	manufacturingComplex.SetBuilder(busBuilder)
	manufacturingComplex.Construct()

	bus := busBuilder.GetVehicle()
	if bus.Wheels != 8 {
		t.Errorf("Wheels on a bus must be 8 and they were %d\n", bus.Wheels)
	}
	if bus.Structure != "Bus" {
		t.Errorf("Structure on a bus must be 'Bus' and was %s\n", bus.Structure)
	}
	if bus.Seats != 30 {
		t.Errorf("Seats on a bike must be 30 and they were %d\n", bus.Seats)
	}
}
