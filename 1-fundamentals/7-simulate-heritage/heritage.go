package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
	height   float32
}
type student struct {
	person     // embedding person struct
	course     string
	university string
}
type teacher struct {
	person            // embedding person struct
	subject           string
	yearsOfExperience int
}

func main() {
	// Creating a person instance
	person1 := person{
		name:     "Alice",
		lastName: "Johnson",
		age:      28,
		height:   5.4,
	}
	// Creating a student instance
	student1 := student{
		person: person{
			name:     "John",
			lastName: "Doe",
			age:      20,
			height:   5.9,
		},
		course:     "Computer Science",
		university: "XYZ University",
	}

	// Creating a teacher instance
	teacher1 := teacher{
		person: person{
			name:     "Jane",
			lastName: "Smith",
			age:      35,
			height:   5.6,
		},
		subject:           "Mathematics",
		yearsOfExperience: 10,
	}

	fmt.Println("Person:", person1)
	fmt.Println("Student:", student1)
	fmt.Println("Teacher:", teacher1)
}
