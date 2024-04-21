package main

import "fmt"

func calcularMedia(nota1, nota2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para calcular média")

	media := (nota1 + nota2) /2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(calcularMedia(5, 6))
}