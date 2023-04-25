package main

import "fmt"

// Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.
func main() {
	numeros := []float64{1, 2, 3, 4}

	produto := 1.0

	for _, num := range numeros {
		produto *= num
	}
	fmt.Println(produto)
}
