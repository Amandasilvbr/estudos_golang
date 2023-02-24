/*
	Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
		- Nome
		- Sobrenome
		- Sabores favoritos de sorvete
	Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sorvete   []string
}

func main() {
	pessoa_1 := pessoa{
		nome:      "Márcia",
		sobrenome: "Silva",
		sorvete:   []string{"baunilha", "morango"}}

	pessoa_2 := pessoa{
		nome: "Roberta",
		sobrenome:"Silva",
		sorvete: []string{"chocolate","napolitano"}}
	
	
	for _, range_sorvetes1 := range pessoa_1.sorvete{
		fmt.Println(range_sorvetes1)
	}

	for _, range_sorvetes2 := range pessoa_2.sorvete{
		fmt.Println(range_sorvetes2)
	}
}
