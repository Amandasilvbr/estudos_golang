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
	agrupamento_de_dados()

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
	default:
		fmt.Println("o número não é igual a 5 e nem menor que 5")
	}
}

func agrupamento_de_dados() {

	var x [4]int

	/* Agrupamento de dados
	Array
		-Arrays são agrupamentos de dados com índices próprios
		-x[4] quer dizer que o o array possui 4 valores, começando do índice 0 até o 3, (0,1,2,3), ou seja, 4 valores contando com o 0
		-O método len pode ser utilizado para descobrir quantos valores o array tem
		-Não é possível mudar o tamanho do array
	*/

	fmt.Println(len(x))

	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	/* Agrupamento de dados
	Slice
		-Diferente do array que é criado com um número definido de valores, o slice não precisa ter essa informação
		-É possível mudar o tamanho do slice
		-A função range percorre o slice inteiro
	*/

	z := []int{1, 2, 3, 4, 5}
	fmt.Println(z)

	frutas := []string{"Maça", "Banana", "Laranja", "Melancia"}

	for i, range_frutas := range frutas {
		fmt.Println("A fruta da posição", i, "é:", range_frutas)
	}

	/* Agrupamento de dados
	Slice
		-Para pegar de um determinado índice até outro:
	*/

	pizzas := []string{"Mussarela", "Abacaxi", "Quatro queijos", "Moda"}
	fmt.Println(pizzas[1:4])

	/* Agrupamento de dados
	Slice
		-Para anexar a uma slice, utiliza-se o append
		-Se precisar anexar uma slice a outra, precisa utilizar o operador "...", para que cada item da slice seja integrado 
	*/

	numeros := []int{1, 2, 3, 4}

	numeros = append(numeros, 10, 11, 12)

	fmt.Println(numeros)

	slice_1 := []int{10,20,30,30}
	slice_2:= []int{1,2,3,4}
	slice_3 := append(slice_1, slice_2...)

	fmt.Println(slice_3)

	
	/* Agrupamento de dados
	Slice- Make
		-O make cria uma slice de x elementos, porém com capacidade de y elementos, assim como no exemplo abaixo. N = 5, Capacidade = 10 
		-Melhor para a performance do código
		-Uma slice multidimensional é uma slice que contém outro slice dentro slice_1[x][y]
	*/

	slice_4 := make([]string, 5, 10)
	slice_4[0], slice_4[1] = "oi","tchau"
	fmt.Println(slice_4)

	/* Agrupamento de dados
	Slice- Maps
		-O make cria uma slice de x elementos, porém com capacidade de y elementos, assim como no exemplo abaixo. N = 5, Capacidade = 10 
		-Melhor para a performance do código
		-Uma slice multidimensional é uma slice que contém outro slice dentro slice_1[x][y]
	*/

	produto := map[string]int{"Calça": 50, "Camisa": 40}
	fmt.Println(produto)
	fmt.Println(produto["Calça"])
}
