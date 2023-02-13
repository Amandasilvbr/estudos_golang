package main

import(
	"fmt"
)

var x int
var y string
var z bool

func main(){
	
	fmt.Println("Exercício 2")

	/*
		Use var para declarar variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. 
		Utilize os seguintes identificadores e tipos para essas variáveis:

		1- x será int
		2- y será string
		3- z será bool

		Depois, demonstre os valores de cada identificador. 
		O compilador atribuiu valores para estas variáveis. Como esses valores se chamam?

		Resposta: são valores 0, pois as variáveis não foram iniciadas 

	*/

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}