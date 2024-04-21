package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i", i)
		time.Sleep(time.Second)
	}

	for j:= 0; j < 10; j++{
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Second)
	} 

	nomes := [4]string{"Lucas", "João", "Laryssa", "Marcos"}

	for indice, nome := range nomes {
		fmt.Println(indice,nome)
	}

	user := map[string] string {
		"nome:": "Walter",
		"sobrenome:": "White",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}	
}