package main

import (
	"fmt"
	"sync"
	"time"
)

func write(txt string, t int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 20; i++ {
			c <- fmt.Sprintf("Receiver value: %s id: %d", txt, i)
			time.Sleep(time.Duration(t) * time.Millisecond)
		}
		close(c) // Fecha o canal ao final
	}()
	return c
}

func main() {
	chanel1 := multiplexer(write("Channel 1", 250), write("Channel 2", 500), write("Channel 3", 1000))
	for msg := range chanel1 {
		fmt.Println(msg)
	}
	fmt.Println("All messages have been processed.")
}

func multiplexer(channels ...<-chan string) <-chan string {
	c := make(chan string)
	var wg sync.WaitGroup

	// Adiciona uma goroutine ao WaitGroup para cada canal
	wg.Add(len(channels))
	for _, ch := range channels {
		go func(ch <-chan string) {
			defer wg.Done() // Marca a goroutine como concluída
			for msg := range ch {
				c <- msg
			}
		}(ch)
	}

	// Fecha o canal `c` quando todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}
