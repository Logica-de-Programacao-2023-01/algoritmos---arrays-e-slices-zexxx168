package main

import "fmt"

// Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que deverão ser trocados de posição. Imprima o Slice resultante.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(slice)

	var num1, num2 int
	fmt.Println("digite o primeiro eo segundo indice a ser trocado")
	fmt.Scan(&num1, &num2)

	if num1 < 0 || num1 >= len(slice) || num2 < 0 || num2 >= len(slice) {
		fmt.Println("indices invalidos.")
		return
	}
	slice[num1], slice[num2] = slice[num2], slice[num1]

	fmt.Println(slice)
}
