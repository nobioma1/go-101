package main

import "fmt"

func main() {
	// create an array of names
	fruits := [10]string{
		"mango",
		"pear",
		"orange",
		"grape",
		"pineapple",
		"apple",
		"banana",
		"blueberry",
		"cashew",
		"raspberry",
	}

	// LOOP TILL CONDITION
	var index int
	var str string
	for index < 10 {
		str += (fruits[index] + " ")
		index++
	}
	fmt.Println(str)

	// LOOP TILL CONDITION WITH POST INCREMENT CLAUSE
	for i := 0; i < 5; i++ {
		fmt.Println(i+1, fruits[i])
	}
	fmt.Println(str)

	// INFINITE LOOP
	var counter int
	for {
		if fruits[counter][4:] == "berry" {
			fmt.Println(fruits[counter])
			break
		}
		counter++
	}

	// LOOPING WITH range
	for i := range fruits {
		fmt.Println(fruits[i], i)
	}

	contacts := map[string]int{
		"tuesday":   3,
		"wednesday": 4,
		"thursday":  5,
		"friday":    6,
	}

	// LOOPING THROUGH MAPS
	for k, v := range contacts {
		fmt.Printf("Key: %s, value %v\n", k, v)
	}
}
