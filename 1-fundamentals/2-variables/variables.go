package main

import (
	"fmt"
)

func main() {
	var variable1 string = "Variable 1"
	variable2 := "Variable 2"

	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "Variable 3"
		variable4 string = "Variable 4"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "Variable 5", "Variable 6"
	fmt.Println(variable5, variable6)

	const const1 string = "Const 1"
	fmt.Println(const1)

	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}
