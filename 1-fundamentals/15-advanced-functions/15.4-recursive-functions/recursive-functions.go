package main

import "fmt"

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("Fibonacci")

	count := uint(40)
	for i := uint(1); i < count; i++ {
		fmt.Println(fibonacci(i))
	}
}
