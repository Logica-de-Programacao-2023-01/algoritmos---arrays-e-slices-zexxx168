package main

import "fmt"

// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
func main() {
	var larg int
	fmt.Println("digite a largura da lista")
	fmt.Scan(&larg)

	soma := 0

	lista := make([]int, larg)

	fmt.Println("digite os numero da lista")
	for i := 0; i < larg; i++ {
		fmt.Scan(&lista[i])
		soma += lista[i]
	}
	fmt.Println(soma)
}
