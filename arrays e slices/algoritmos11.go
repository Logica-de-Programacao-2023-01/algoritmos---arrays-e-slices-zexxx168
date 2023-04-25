package main

//Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.
import "fmt"

func main() {
	matriz := [2][3]int{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("informe o valor da coluna [%d][%d]: ", i, j)
			fmt.Scan(&matriz[i][j])

		}
	}
	fmt.Printf("o valor da sua matriz bidimensional e:")
	for i := 0; i < 2; i++ {
		fmt.Println()
		for j := 0; j < 3; j++ {
			fmt.Printf("%d", matriz[i][j])
		}
	}
	fmt.Println()
	var linha, coluna int

	fmt.Println("digite o valor de linha")
	fmt.Scan(&linha)
	fmt.Println("digite o valor de coluna")
	fmt.Scan(&coluna)

	fmt.Println("o valor do numero e: ", matriz[linha][coluna])
}
