package main

import "fmt"

func calculos(num1, num2 int) (soma int, subtracao int) {
	soma = num1 + num2
	subtracao = num1 - num2
	return
}

func main() {
	soma, subtacao := calculos(10, 20)
	fmt.Println(soma, subtacao)
}