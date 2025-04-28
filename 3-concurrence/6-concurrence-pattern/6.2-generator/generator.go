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
	chanel1 := write("Channel 1", 250)
	chanel2 := write("Channel 2", 500)
	chanel3 := write("Channel 3", 1000)

	var wg sync.WaitGroup
	wg.Add(3) // Adiciona 3 goroutines ao WaitGroup

	// Função para consumir e imprimir os valores de um canal
	printChannel := func(ch <-chan string) {
		defer wg.Done() // Marca como concluído ao final
		for msg := range ch {
			fmt.Println(msg)
		}
	}

	// Inicia uma goroutine para cada canal
	go printChannel(chanel1)
	go printChannel(chanel2)
	go printChannel(chanel3)

	wg.Wait() // Aguarda todas as goroutines finalizarem
	fmt.Println("All channels have been processed.")
	fmt.Println("End of program.")
}
