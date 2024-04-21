package main

import "fmt"

func main() {
	fmt.Print("Ponteiros\n")

	var num1 int = 10
	var num2 int = num1
	fmt.Println(num1, num2)

	num1++
	fmt.Println(num1, num2)
	
	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var num3 int
	var ponteiro *int

	num3 = 100
	ponteiro = &num3

	fmt.Println(num3, *ponteiro) // desreferenciação

	num3 = 150

	fmt.Println(num3, *ponteiro) // desreferenciação
}