package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	numeros = append(numeros[:2], numeros[3:]...)

	fmt.Println(numeros)
}
