package main

import (
	"fmt"

	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("write from main package")
	assistant.Assistant()

	// checkmail package example
	valid := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(valid)

	// checkmail package example
	invalid := checkmail.ValidateFormat("not is a email")
	fmt.Println(invalid)
}
