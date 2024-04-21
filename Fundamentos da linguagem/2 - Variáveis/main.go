package main

import "fmt"

func main(){
	x := "var 1"
	fmt.Println(x)

	y, z := 42, 76
	fmt.Println(y, z)

	y, z = z, y
	fmt.Println(y, z)
}