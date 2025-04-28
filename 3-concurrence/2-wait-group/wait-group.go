package main

import (
	"fmt"
	"sync"
	"time"
)

func write(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt, i)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(3) // Add two goroutines to the wait group

	go func() {
		write("Hello")
		waitGroup.Done() // Mark this goroutine as done
	}()

	go func() {
		write("World")
		waitGroup.Done() // Mark this goroutine as done
	}()
	go func() {
		write("!")
		waitGroup.Done() // Mark this goroutine as done
	}()
	waitGroup.Wait() // Wait for all goroutines to finish
	fmt.Println("All goroutines finished executing.")
}
