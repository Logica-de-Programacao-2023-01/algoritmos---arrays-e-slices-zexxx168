package main

import "fmt"

// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.
func main() {
	numeros := [3][2]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("digite o valor da posicao [%d][%d]:", i, j)
			fmt.Scan(&numeros[i][j])
		}
	}
	fmt.Println("matriz resultante e:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d", numeros[i][j])
		}
		fmt.Println()
	}
}
