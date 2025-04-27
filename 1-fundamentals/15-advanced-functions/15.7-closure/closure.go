package main

import "fmt"

func closure() func() {
	txt := "Into the closure function"

	function := func() {
		fmt.Println(txt)
	}
	return function
}

func main() {
	txt := "Into the main function"
	fmt.Println(txt)
	closureFunction := closure()
	closureFunction()
}
