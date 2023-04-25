package main

import "fmt"

func main() {
	// Cria um Slice de inteiros com tamanho 8
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// Solicita ao usuário que informe os índices de elementos a serem trocados
	var i, j int
	fmt.Println("Informe os índices de elementos a serem trocados:")
	fmt.Scan(&i, &j)

	// Verifica se os índices informados são válidos
	if i < 0 || i >= len(s) || j < 0 || j >= len(s) {
		fmt.Println("Índices inválidos!")
		return
	}

	// Troca os elementos de posição
	s[i], s[j] = s[j], s[i]

	// Imprime o Slice resultante
	fmt.Println(s)
}
