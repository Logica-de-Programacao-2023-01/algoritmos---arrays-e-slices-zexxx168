package main

import "fmt"

func main() { //Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
	var numeros [3]int
	var soma int

	fmt.Println("digite um numero")

	for i := 0; i < 3; i++ {
		fmt.Scan(&numeros[i])
		soma = soma + numeros[i]
	}

	fmt.Println(soma)

}
