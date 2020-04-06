package main

import "fmt"

func main() {
	// explicit variable declaration
	var port int32
	port = 3000
	fmt.Printf("Current port: %d \n", port)

	// One line declaration
	var speed float32 = 358.69
	fmt.Printf("It's travelling at %.2f knots\n", speed)

	// inferred type variable declaration
	age := 20
	fmt.Println(age)

	// ./main.go:18:6: no new variables on left side of :=
	// age := 21
	// fmt.Println(age)
}
