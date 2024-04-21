package main

import "fmt"

func recuperarExecucao(){
	if r:= recover(); r != nil {
		fmt.Println("Execução Recuperada com sucesso")
	}
}

func calcularMedia(nota1, nota2 float32) bool {
	defer recuperarExecucao()
	media := (nota1 + nota2) /2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(calcularMedia(6, 6))
}