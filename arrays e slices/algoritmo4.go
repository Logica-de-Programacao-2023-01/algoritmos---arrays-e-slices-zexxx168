package main

import "fmt"

func main() {
	var tamanho int
	var soma int

	fmt.Println("digite quantos numeros vc quer somar")
	fmt.Scan(&tamanho)
	numeros := make([]int, tamanho)
	fmt.Println("digite um numero")
	soma = 0

	for i := 0; i < tamanho; i++ {
		fmt.Scan(&numeros[i])
		soma = soma + numeros[i]
	}
	fmt.Println(soma)
}
