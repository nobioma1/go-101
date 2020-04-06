package main

import "fmt"

// Arrays have a fixed length and are used to stores items of the same type

func main() {
	// verbose declaration of an array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	fmt.Println(arr)

	// a more elaborate definition
	words := [3]string{"apple", "bapple", "capple"}
	fmt.Println(words)
}
