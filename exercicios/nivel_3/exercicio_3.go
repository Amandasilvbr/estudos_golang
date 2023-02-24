package main

import (
	"fmt"
)

func main(){
	/*
	Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
		- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
		- Do quinto ao último item do slice (incluindo o último item!)
		- Do segundo ao sétimo item do slice (incluindo o sétimo item!)
		- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
		- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	*/

	valores := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(valores[:3])
	fmt.Println(valores[2:7])
	fmt.Println(valores[5:])
	fmt.Println(valores[3:9])
	fmt.Println(valores[3:(len(valores)-1)])

}