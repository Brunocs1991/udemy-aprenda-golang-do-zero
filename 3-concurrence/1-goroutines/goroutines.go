package main

import (
	"fmt"
	"time"
)

func write(txt string) {
	for i := 0; i < 100; i++ {
		fmt.Println(txt, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

func main() {
	go write("Parallelism different from concurrency 1")
	write("Parallelism different from concurrency 2")
	fmt.Println("Done")
	// Wait for goroutines to finish (not a good practice, just for demonstration)
	time.Sleep(2 * time.Second)
	fmt.Println("Exiting main function")
}
