package main 

import(
	"fmt"
)

func main(){
	/* 
	Usando uma literal composta:
		- Crie uma slice de tipo int
		- Atribua 10 valores a ela
		- Utilize range para demonstrar todos estes valores.
		- E utilize format printing para demonstrar seu tipo.
	*/

	numeros := []int{10,20,30,40,50,60,70,80,90,100}

	for _, num := range numeros{
		fmt.Println(num)
	}

	fmt.Printf("%T",numeros)

}