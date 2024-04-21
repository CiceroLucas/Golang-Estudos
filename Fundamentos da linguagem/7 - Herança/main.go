package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade  	  uint8
	altura    uint8
}

type aluno struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"João", "Arthur", 23, 167}

	e1 := aluno{p1, "PSI", "UNIL"}

	fmt.Println(e1)
}