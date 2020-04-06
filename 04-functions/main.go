package main

import "fmt"

func main() {
	// Multiple assignment
	age, class1 := classifyAge(20)
	fmt.Println(age, class1)
	// ignore declaration
	_, class2 := classifyAge(50)
	fmt.Println(class2)

}

// function prints 'Hello World'
func printHelloWorld() {
	fmt.Println("Hello World")
}

// takes an string argument and returns a string
func log(message string) string {
	return message
}

// type and return two return values
func classifyAge(age int) (int, bool) {
	return age, age > 50
}

// parameters with the same type can share the same type definition
func sum(a, b int) int {
	return a + b
}
