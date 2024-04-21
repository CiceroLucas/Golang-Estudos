package main

import "fmt"

func main() {
	// OPERADORES ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 5
	multiplicacao := 2 * 5
	resto := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	// OPERADORES DE ATRIBUIÇÃO
	var txt string = "string"
	txt2 := "string"
	fmt.Println(txt, txt2)

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// OPERADORES UNÁRIOS
	numero := 10
	numero ++
	numero += 15
	fmt.Println(numero)

	numero --
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
}