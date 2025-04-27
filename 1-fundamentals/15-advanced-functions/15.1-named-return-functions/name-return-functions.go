package main

import "fmt"

func matematicalOperations(n1,n2 int) (sum, sub int){
	// Named return function
	sum = n1 + n2
	sub = n1 - n2
	return // Return the named variables sum and sub
	
}

func main() {
	fmt.Println("Named return functions")
	n1 := 10
	n2 := 5
	sum, sub := matematicalOperations(n1, n2)
	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", sub)
}