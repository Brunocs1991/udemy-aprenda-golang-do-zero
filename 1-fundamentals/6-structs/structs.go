package main

import "fmt"

// Person struct represents a person with a name and age.
type person struct {
	Name    string  // Name of the person
	Age     int     // Age of the person
	address address // Address of the person (embedded struct)
}

type address struct {
	Street string // Street address
	number string // House number (unexported field)
}

func main() {
	// Create a new Person instance
	person1 := person{
		Name: "Alice",
		Age:  30,
		address: address{
			Street: "123 Main St",
			number: "1A", // This field is unexported and cannot be accessed outside the package
		},
	}

	fmt.Println(person1) // Output: {Alice 30}

	var person2 person   // Declare a variable of type Person
	fmt.Println(person2) // Output: { 0} (default values for string and int types)
	// Assign values to the fields of person2
	person2.Name = "Bob"
	person2.Age = 25     // Create an address instance
	fmt.Println(person2) // Output: {Bob 25}

	person3 := person{Name: "Charlie"} // Age will be 0 by default
	fmt.Println(person3)               // Output: {Charlie 0}

	person4 := person{Age: 40} // Name will be empty by default
	fmt.Println(person4)       // Output: { 40} (default value for string is empty)
}
