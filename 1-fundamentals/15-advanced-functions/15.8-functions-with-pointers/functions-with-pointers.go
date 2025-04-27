package main

import "fmt"

func signalInvert(number int) int {
	return number * -1
}

func signalInvertPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 20

	invertNumber := signalInvert(number)
	fmt.Println("The number is: ", invertNumber)


	newNumber := 40
	fmt.Println("The number is: ", newNumber)
	signalInvertPointer(&newNumber)
	fmt.Println("The number is: ", newNumber)
}
