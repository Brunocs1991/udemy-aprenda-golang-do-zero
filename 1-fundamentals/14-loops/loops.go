package main

import (
	"fmt"
	"time"
)

func main() {
	// While loop (using for)
	j := 0
	for j < 10 {
		fmt.Println("While loop iteration:", j)
		time.Sleep(1 * time.Second) // Sleep for 1 second
		j++
	}

	// For loop
	for i := 0; i < 10; i+=2 {
		fmt.Println("For loop iteration:", i)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}

	// for range loop
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("For range loop iteration %d: %s\n", index, value)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}

	for index, char := range "hello" {
		fmt.Printf("For range loop iteration %d: %c\n", index, char)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}

	// for in map loop
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	for key, value := range m {
		fmt.Printf("For in map loop: %s -> %d\n", key, value)
		time.Sleep(1 * time.Second) // Sleep for 1 second
	}

	// Infinite loop
	for {
		fmt.Println("Infinite loop iteration")
		time.Sleep(1 * time.Second) // Sleep for 1 second
		break // Break out of the infinite loop after one iteration
	}
	
}
