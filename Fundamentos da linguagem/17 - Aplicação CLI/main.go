package main

import (
	"cli/app"
	"log"
	"os"
)

func main() {
	aplicação := app.Gerar()
	if erro := aplicação.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
	
}