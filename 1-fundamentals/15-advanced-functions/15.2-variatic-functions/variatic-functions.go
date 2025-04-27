package main

import "fmt"

// Variadic function that takes a variable number of int arguments and returns their sum
func sum(numbers ...int) int {
	fmt.Println(numbers)
	// The ... operator indicates that the function can take a variable number of arguments
	// The type of the argument is int
	// The function returns an int
	// The function takes a variable number of int arguments and returns their sum
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func write(txt string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(txt, number)
	}
}

func main() {
	// Variadic function
	// The ... operator indicates that the function can take a variable number of arguments
	// The type of the argument is int
	// The function returns an int
	// The function takes a variable number of int arguments and returns their sum
	sum1 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Sum:", sum1)

	sum2 := sum()
	fmt.Println("Sum:", sum2)

	write("Number:", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
