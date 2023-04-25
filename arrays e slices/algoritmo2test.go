package main

//Crie um Slice de inteiros com os valores 1, 2, 3, 4 e 5. Remova o terceiro elemento e imprima o Slice resultante.
import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}

	numeros2 := append(numeros[:2], numeros[3:]...)

	fmt.Println(numeros2)

}
