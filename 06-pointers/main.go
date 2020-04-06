package main

import "fmt"

// User struct
type User struct {
	Name   string
	Amount float32
}

func main() {
	// create a user struct value
	tunde := User{
		Name:   "Tunde Ojo",
		Amount: 10,
	}

	fmt.Println(tunde) // {Tunde Ojo 10}
	add(tunde, 2000)   // ?? should add 2000 to tunde's amount
	fmt.Println(tunde) // {Tunde Ojo 10}

	// Arguments are passed to functions as copies, so the use of pointers is needed.

	// Using pointers
	// * - dereference
	// & - addressof

	john := &User{
		Name:   "John Doe",
		Amount: 100,
	}

	fmt.Println(*john) // {John Doe 100}
	addALot(john)
	fmt.Println(*john) // {John Doe 100600}
}

func add(user User, amount float32) {
	user.Amount += amount
}

func addALot(user *User) {
	user.Amount += 100500
}
