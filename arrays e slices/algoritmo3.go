package main

//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.
import "fmt"

func main() {
	var numeros [4]float64
	soma := 1.0

	fmt.Println("digite um numero")

	for i := 0; i < 4; i++ {
		fmt.Scan(&numeros[i])
		soma = soma * numeros[i]
	}
	fmt.Println(soma)
}
