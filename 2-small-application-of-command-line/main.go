package main

import (
	"log"
	"os"

	"small-application-of-command-line/app"
)

func main() {
	application := app.Gerar()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
