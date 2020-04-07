package main

import "fmt"

// A slice is a lightweight structure that wraps and represents a portion of an array.

func main() {
	// Declaring and Initializing a slice
	// method 1
	numbers := []int{1, 2, 3, 4, 5, 6}
	numbers = append(numbers, 11, 12, 13)
	fmt.Println(numbers)

	// method 2
	// make([]type, length, capacity)

	letters := make([]string, 0, 6)
	letters = append(letters, "a", "b", "c", "d", "e", "f")
	fmt.Println(len(letters))
	letters[3] = "E"
	fmt.Println(letters)
	letters = append(letters, "z")
	fmt.Println(letters)

	aSlice := letters[0:3]
	fmt.Println(aSlice)

}
