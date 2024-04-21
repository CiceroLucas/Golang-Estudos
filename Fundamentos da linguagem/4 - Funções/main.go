package main

import "fmt"

func somar(num1 int8, num2 int8) int8 {
	return num1 + num2

}

func calculos(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	sub := num1 - num2

	return soma, sub
}

func main() {
	resSoma := somar(10, 20)
	fmt.Println(resSoma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	res := f("Chamando função de texto")
	fmt.Println(res)

	_, sub := calculos(10, 15)
	fmt.Println(sub)

}