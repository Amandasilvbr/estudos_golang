package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	fmt.Println("Exercício 3")

	/*
				- Utilizando a solução do exercício anterior:
		    	1- Em package-level scope, atribua os seguintes valores às variáveis a seguir:
		        	-Para "x" atribua 42
		        	-Para "y" atribua "James Bond"
		        	-Para "z" atribua true

		    	2. Na função main:
		        	-Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
		        	-Demonstre a variável "s".

	*/

	s :=  fmt.Sprintf("%v %v %v",x, y,z)
	fmt.Print(s)

}
