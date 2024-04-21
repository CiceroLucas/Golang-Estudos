package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Numero invalido"
	}
}

func diaDaSemana2(num int) string {
	var dias string

	switch num {
	case 1:
		dias = "Domingo"
	case 2:
		dias = "Segunda-feira"
	case 3:
		dias = "Terça-feira"
	case 4:
		dias = "Quarta-feira"
	case 5:
		dias = "Quinta-feira"
	case 6:
		dias = "Sexta-feira"
	case 7:
		dias = "Sábado"
	default:
		dias = "Numero invalido"
	}
	return dias
}

func main() {
	dia := diaDaSemana(9)
	fmt.Println(dia)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}