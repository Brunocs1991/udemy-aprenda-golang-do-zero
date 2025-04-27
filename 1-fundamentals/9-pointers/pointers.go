package main

import "fmt"

func main() {
	var variable1 int = 10
	var variable2 int = variable1

	fmt.Println(variable1, variable2) // Output: 10 10
	variable2++
	fmt.Println(variable1, variable2) // Output: 10 11

	//Ponters is a memory address of a variable
	var variable3 int
	var pointer *int // pointer is a pointer to variable3
	fmt.Println(variable3, pointer) // Output: 0 <nil>
	variable3 = 100
	pointer = &variable3 // pointer now points to the address of variable3
	fmt.Println(variable3, pointer) // Output: 0 0xc00000a0b8
	fmt.Println(*pointer) // Output: 100
	*pointer = 200 // change the value of variable3 using pointer
	fmt.Println(variable3, pointer) // Output: 200 0xc00000a0b8
	fmt.Println(*pointer) // Output: 200

}