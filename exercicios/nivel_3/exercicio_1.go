package main

import (
	"fmt"
)

func main() {
	/* 
	Usando uma literal composta:
	    - Crie um array que suporte 5 valores to tipo int
	    - Atribua valores aos seus índices
		- Utilize range e demonstre os valores do array.
		- Utilizando format printing, demonstre o tipo do array.
	*/

	valores := [5]int{1,2,3,4,5}

	for _, range_valores := range valores{
		fmt.Println(range_valores)
	}

	fmt.Printf("%T", valores)
}
