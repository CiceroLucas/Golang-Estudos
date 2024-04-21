package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := uint(7)
	fmt.Printf("Posição: %v\n",fibonacci(posicao))

	fmt.Println("[Sequência de Fibonacci]")

	for i := uint(0); i < posicao; i++{
		fmt.Println(fibonacci(i))
	}
}