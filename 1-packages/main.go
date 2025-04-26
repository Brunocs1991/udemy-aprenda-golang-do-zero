package main

import (
	"module/assistant"
	"fmt"
)

func main() {
	fmt.Println("write from main package")
	assistant.Assistant()
}