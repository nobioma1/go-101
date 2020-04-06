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
}
