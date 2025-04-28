package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for i := 0; i < 50; i++ {
			channel1 <- "channel 1"
			time.Sleep(500 * time.Millisecond)
		}
		close(channel1) // Fecha o canal após o envio
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel2 <- "channel 2"
			time.Sleep(2 * time.Second)
		}
		close(channel2) // Fecha o canal após o envio
	}()

	for {
		select {
		case message1, ok := <-channel1:
			if !ok { // Verifica se o canal foi fechado
				channel1 = nil // Evita bloqueio no select
			} else {
				fmt.Println("Received from channel 1:", message1)
			}
		case message2, ok := <-channel2:
			if !ok { // Verifica se o canal foi fechado
				channel2 = nil // Evita bloqueio no select
			} else {
				fmt.Println("Received from channel 2:", message2)
			}
		}

		// Sai do loop quando ambos os canais forem fechados
		if channel1 == nil && channel2 == nil {
			break
		}
	}
	fmt.Println("All channels closed. Exiting.")
	fmt.Println("End of program.")
}
