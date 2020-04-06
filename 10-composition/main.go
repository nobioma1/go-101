package main

import "fmt"

// COMPOSITION: This is the act of including a structure into another

type Person struct {
	Name string
	age  int
}

func (person *Person) Walk() string {
	return person.Name + " is walking..."
}

type Teacher struct {
	*Person
	Class string
}

func (teacher *Teacher) Teach() string {
	return teacher.Name + " is teaching..."
}

func main() {
	teacher := &Teacher{
		Person: &Person{
			Name: "Friday Monday",
			age:  32,
		},
		Class: "Jss 2",
	}

	fmt.Println(teacher.Walk())
}
