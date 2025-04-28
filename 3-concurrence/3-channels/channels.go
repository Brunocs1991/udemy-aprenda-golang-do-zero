package main

import (
	"fmt"
	"time"
)

func write(txt string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- txt
		time.Sleep(time.Second * 2)
	}
	close(channel) // Close the channel when done to avoid deadlocks
}

func main() {
	channel := make(chan string)
	go write("Hello world !", channel)
	for msg := range channel {
		fmt.Println(msg)
	}
	fmt.Println("End of program")
}
