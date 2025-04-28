package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "Hello World!"       // This will cause a deadlock because the channel is unbuffered and there is no receiver
	channel <- "Programming in Go!" // This will cause a deadlock because the channel is unbuffered and there is no receiver

	message := <-channel  // This line will never be reached because the program will deadlock before it gets here
	message2 := <-channel // This line will never be reached because the program will deadlock before it gets here
	fmt.Println(message)  // This line will never be reached because the program will deadlock before it gets here
	fmt.Println(message2) // This line will never be reached because the program will deadlock before it gets here
}
