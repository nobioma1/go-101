package main

import "fmt"

// Person struct
type Person struct {
	Name   string
	Age    int
	Height float32
}

func main() {
	noble := Person{
		Name: "Noble Obioma",
		Age:  22,
	}

	noble.Height = 6.2
	fmt.Println(noble)
	fmt.Println(noble.Age)

	fin := Person{}
	fin.Name = "Fin"
	fin.Age = 15
	fin.Height = 5.2
	fmt.Println(fin)
	fmt.Println(fin.Height)

}
