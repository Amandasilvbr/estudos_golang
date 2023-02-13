package main

import (
	"fmt"
)

// escopo global, variável acessível dentro do código inteiro

func main() {

	// escopo local, variável disponível apenas dentro desse code block

	fmt.Println("\n", "Exemplo 1")
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

	// Criação de tipos: consultar exercício 4

	/* Tipo booleano, operadores relacionais e operadores condicionais em go

		Em go:

		Igualdade: ==
		Diferença: !=
		Maior/ menor: < >

	Normalmente os operadores relacionais devolvem um valor booleano (verdadeiro ou falso)

		Exemplo:
		100 > 200 ? false
		200 > 100 ? true

	Os operadores condicionais também retornam verdadeiro ou false
		&&: Operador e, ou seja, se ambas as condições forem verdadeiras
		||: Operador ou, ou seja, se uma das condições forem verdadeiras, não necessariamente as duas

	*/
	loops_condicionais()

}

func loops_condicionais() {

	/* Loops
	Inicia a variável iterável (x), coloca a condição de saída do laço, itera o valor x com +1
		-O código vai imprimir 10 vezes os números e depois parar, pois essa foi a condição estabelecida
		-Para parar o loop, é possível usar o break e definir um if que para o loop a partir de x condição
		-Para pular para o próximo laço, é possível usar o continue e pular para o próximo loop sem rodar o código inteiro
	Em go, o loop while não existe
	*/

	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}

	/* Condicionais
	If
		-if (se) condição é verdadeira, faz o que está no código dentro desse if
		-else, se não for verdadeiro, faz o que está no código dentro do else
		-se tiver mais que uma condição usar o else if e a condição e depois um else simples

	*/
	y := 8
	if y < 10 {
		fmt.Println("O número é menor que 10")
	} else {
		fmt.Println("O número é maior que 10")
	}

	/* Condicionais
	Switch
		-Testa os casos e faz o que o código manda se a corresponência for verdadeira
		-Ele testa caso a caso e faz apenas o que é verdadeiro na primeira condição estabelecida
		-Default é o que vai ser feito se nenhuma condição do case for atendida
	*/

	z := 10

	switch {
	case z == 5:
		fmt.Println("O número é igual a 5")
	case z <= 5:
		fmt.Println("o número é menor que 5")
	default: fmt.Println("o número não é igual a 5 e nem menor que 5")
	}

	/* Agrupamento de dados
	
	*/
}
