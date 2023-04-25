package main

import "fmt"

// Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array. Imprima o resultado da busca
func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var numerodousuario int
	var encontrado bool
	fmt.Println("digite o numero ")
	fmt.Scan(&numerodousuario)

	for i := 0; i < len(numeros); i++ {
		if numerodousuario == numeros[i] {
			encontrado = true
			break
		} else {
			continue
		}
	}
	if encontrado {
		fmt.Printf("o numero %d foi encontrado na lista de numeros", numerodousuario)
	} else {
		fmt.Printf("o numero %d nao foi encontrado na lista de numeros", numerodousuario)
	}

}
