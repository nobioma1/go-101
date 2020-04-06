package main

// Go has a 'new' keyword to allocate memory for a type

type Person struct {
	Name string
	Age  int
}

func main() {
	person := new(Person) 
	
	// same as
	
	pn := &Person{}

	p := new(Person) 
	p.Name = "Person Name" 
	p.Age = 4003
	
	// or
	
	person2 := &Person { 
		Name: "Jenny Brown", 
		Age: 9000,
	}
}
