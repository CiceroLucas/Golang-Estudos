package main

import "fmt"

func main() {
	fmt.Printf("Maps\n")

	user := map[string]string {
		"nome": "Lucas",
		"sobrenome": "Sousa",
	}

	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Lucas",
			"ultimo": "Sousa",
		},
		"curso": {
			"nome": "ADS",
			"campus": "IFPB-CZ",
		},
	}

	fmt.Println(user2)
	delete(user2, "nome")
	fmt.Println(user2)
}