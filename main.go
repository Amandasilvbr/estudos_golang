package main

import (
	"fmt"
)

// escopo global, variável acessível dentro do código inteiro

func main() {

	// escopo local, variável disponível apenas dentro desse code block

	fmt.Println("\n","Exemplo 1")
	fmt.Println("Hello world!")

	// := usada para a criação de variáveis

	o := 10
	a := "ola"
	b := "tchau"

	// Println: escreve na tela o valor passado com uma linha no final
	fmt.Println(a)
	fmt.Println(o)

	// Print: mesmo que o anterior, sem a linha no final
	fmt.Print(b)
	fmt.Print(o)
	
	// Criação de tipos: exercício 4

}


