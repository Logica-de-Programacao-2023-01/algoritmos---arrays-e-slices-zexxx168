package main

import "fmt"

// Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array.
// uma curiosidade e que tb pd ser feito com range
func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	soma := 0

	for i := 0; i < len(array); i += 2 {
		soma += array[i]
	}

	fmt.Println("Soma dos elementos nas posições pares:", soma)

}
