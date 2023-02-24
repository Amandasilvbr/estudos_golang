/* 
	Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
	- Demonstre os valores do map utilizando range.
	- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
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
		sobrenome:"Cardoso",
		sorvete: []string{"chocolate","napolitano"}}
	
	
	for _, range_sorvetes1 := range pessoa_1.sorvete{
		fmt.Println(range_sorvetes1)
	}

	for _, range_sorvetes2 := range pessoa_2.sorvete{
		fmt.Println(range_sorvetes2)
	}

	map_1 := map[string]string{pessoa_1.sobrenome: pessoa_1.nome, pessoa_2.sobrenome: pessoa_2.nome}
	fmt.Println(map_1)

	for _, range_map := range map_1{
		fmt.Println(range_map)
		for _, range_s1 := range pessoa_1.sorvete{
			fmt.Println(range_s1)
		}
		for _, range_s2 := range pessoa_2.sorvete{
			fmt.Println(range_s2)
		}
	}
}