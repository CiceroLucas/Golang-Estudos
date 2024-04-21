package main

import "fmt"

func main() {
	num := 10

	if num > 5 {
		fmt.Println("Maior que 15")
	}

	if outroNum := num; outroNum > 0 {
		fmt.Println("Numero é maior que zero")
	}
}