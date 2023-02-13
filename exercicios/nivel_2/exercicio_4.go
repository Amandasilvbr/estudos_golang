package main

import (
	"fmt"
)

func main() {
	//  Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
	esporteFavorito := "Natação"
	switch {
	case esporteFavorito == "Natação":
		fmt.Println("Natação é legal")
	case esporteFavorito == "Futebol":
		fmt.Println("Futebol é legal")
	}
}
