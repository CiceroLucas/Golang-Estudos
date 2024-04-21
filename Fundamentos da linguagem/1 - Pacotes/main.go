package main

import (
	"fmt"
	"module/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do main")
	assistant.Write()
	error := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(error)
}