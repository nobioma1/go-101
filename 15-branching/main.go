package main

import "fmt"

func main() {
	height(3)
	height(6)
	height(15)

	str := "*+-$"

	for i := range str {
		// can't do string indexing str[i] return 100, 42, 43, 45, 36 of type byte
		decode(str[i : i+1])
	}
}

func height(v int) {
	if v < 5 {
		fmt.Println("short")
	} else if v > 5 && v < 10 {
		fmt.Println("average")
	} else {
		fmt.Println("tall")
	}
}

func decode(s string) {
	switch s {
	case "*":
		fmt.Println("all")
	case "-":
		fmt.Println("less")
	case "+":
		fmt.Println("some")
	default:
		fmt.Println("none")
		panic("Why come here")
	}
}
