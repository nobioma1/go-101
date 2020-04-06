package main

// Structures do not have constructors.
// Instead we create a function that return an instance of the desired type.

type Car struct {
	Make    string
	Model   string
	Mileage float32
}

func main() {

}

// NewCarP using pointers
func NewCarP(make string, model string, mileage float32) *Car {
	return &Car{
		Make:    make,
		Model:   model,
		Mileage: mileage,
	}
}

// NewCar withot pointers
func NewCar(make string, model string, mileage float32) Car {
	return Car{
		Make:    make,
		Model:   model,
		Mileage: mileage,
	}
}
