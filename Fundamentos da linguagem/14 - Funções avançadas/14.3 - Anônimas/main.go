package main

import "fmt"

func main() {
	res := func(txt string) string {
		return fmt.Sprintf("Recebebido -> %s", txt)
	}("Passando parâmetro")

	fmt.Println(res)
}