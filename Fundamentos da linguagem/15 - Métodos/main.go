package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("Usuario %s salvo no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func main() {
	user1 := usuario{"Lucas", 23}
	user1.salvar()
	
	maiorDeIdade := user1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	user2 := usuario{"Levi", 15}
	user2.salvar()
}