package main

import "fmt"

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}

func stundentAproval(n1, n2 float64) bool {
	defer recoverPanic()
	med := (n1 + n2) / 2

	if med > 6 {
		return true
	} else if med < 6 {
		return false
	}
	panic("The average is equal to 6")
}

func main() {
	fmt.Println(stundentAproval(6, 7))
	fmt.Println(stundentAproval(3, 7))
	fmt.Println(stundentAproval(6, 6))
	fmt.Println("after execution")
}
