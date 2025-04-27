package main

import "fmt"

func func1() {
	fmt.Println("Function 1")
}

func func2() {
	fmt.Println("Function 2")
}

func stundentAproval(n1, n2 float64) bool {
	defer fmt.Println("Average calculated, result will be returned")
	fmt.Println("Open the function to validate the student is approved")
	med := (n1 + n2) / 2
	if med >= 6 {
		return true
	}
	return false
}

func main() {
	defer func1()
	func2()
	fmt.Println(stundentAproval(7, 8))
}
