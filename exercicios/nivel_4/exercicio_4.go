/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import(
	"fmt"
)


func main(){
	
structx := struct {
	slice1 []string
	map1 map[int]string
}{
	slice1: []string{"a","b"},
	map1: map[int]string{1:"aa", 2: "bb"},
}

	fmt.Println(structx)
}