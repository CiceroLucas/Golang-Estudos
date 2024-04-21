package main

import (
	"errors"
	"fmt"
)

func main(){
	numero := 1000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//int32 == rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	//byte == uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 12312321321.212
	fmt.Println(numeroReal2)

	var str string = "ola"
	fmt.Println(str)

	str2 := "mundo"
	fmt.Println(str2)

	var texto string
	fmt.Println(texto)

	var boolean bool
	fmt.Println(boolean)

	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}