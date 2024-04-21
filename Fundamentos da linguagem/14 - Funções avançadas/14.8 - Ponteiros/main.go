package main

import "fmt"

func inverterSinal(num int) int {
	return num * -1
}

func inverterSinalComPonteiro(numero *int){
	*numero = *numero * -1
}

func main() {
	num := 20
	numInvertido := inverterSinal(num)
	fmt.Println(numInvertido)

	num2 := 40
	fmt.Println(num2)
	inverterSinalComPonteiro(&num2)
	fmt.Println(num2)

}