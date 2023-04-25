package main

import "fmt"

func main() {
	// Criando um array de 5 elementos
	arr := [5]int{2, 6, 9, 12, 15}

	// Criando um slice vazio para armazenar os múltiplos de 3
	multOfThree := []int{}

	// Iterando pelo array e adicionando os múltiplos de 3 ao slice
	for _, num := range arr {
		if num%3 == 0 {
			multOfThree = append(multOfThree, num)
		}
	}

	// Imprimindo o slice resultante
	fmt.Println(multOfThree)
}
