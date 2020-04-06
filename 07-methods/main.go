package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := &Person{
		Name: "John Doev",
		Age:  20,
	}

	person.Talk()
}

// Talk the ability to tal to the person method
func (person *Person) Talk() {
	fmt.Println(person.Name, " said something")
}
