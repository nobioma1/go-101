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
