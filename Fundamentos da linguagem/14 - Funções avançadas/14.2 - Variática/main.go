package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, value := range numeros {
		total += value
	}

	return total
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)
}