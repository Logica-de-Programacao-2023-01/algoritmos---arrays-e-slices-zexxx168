package main

import (
	"fmt"
)

// Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.
func main() {
	var encontrado bool
	numeros := []int{1, 2, 3, 4, 5}

	var numerodousuario int

	fmt.Println("digite o numero ")
	fmt.Scan(&numerodousuario)

	for _, i := range numeros {
		if i == numerodousuario {
			encontrado = true
			break
		} else {
			encontrado = false
			continue
		}
	}
	fmt.Printf("a sua lista de numeros eh %d", numeros)
	fmt.Println()
	if encontrado {
		fmt.Printf("o numero %d foi encontrado e nao foi adicionado a lista de numeros", numerodousuario)
	} else {
		fmt.Printf("o numero %d nao foi encontrado e foi adicionado a lista de numeros", numerodousuario)
		numeros = append(numeros, numerodousuario)
		fmt.Println()
		fmt.Printf("a sua lista de numeros agora eh %d", numeros)
	}

}
