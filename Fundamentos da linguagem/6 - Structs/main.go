package main

import "fmt"

type user struct {
	nome 		string
	idade 		uint8
	enderero 	endereco
}


type endereco struct {
	logradouro 	string
	numero 		uint8
}

func main() {
	fmt.Println("Arquivo strutcs")

	var newUser user
	
	newUser.nome = "Lucas"
	newUser.idade = 23

	fmt.Println(newUser)

	newEndereco := endereco{"Jose leoncio", 203}

	usuario := user{"Levi", 16, newEndereco}
	fmt.Println(usuario)

	usuario3 := user{nome: "Thiago"}
	fmt.Println(usuario3)
}