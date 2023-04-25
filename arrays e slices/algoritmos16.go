package main

import "fmt"

// Crie um Array de inteiros com 10 elementos. Crie um novo Slice que contenha apenas os elementos pares do Array original. Imprima o novo Slice.
func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	pares := []int{}

	for _, num := range arr {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	fmt.Println(pares)
}
