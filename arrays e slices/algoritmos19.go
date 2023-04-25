package main

import "fmt"

func media(slice []int) float64 {
	soma := 0
	for _, v := range slice {
		soma += v
	}
	return float64(soma) / float64(len(slice))
}

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Printf("Média: %.2f", media(numeros))
}
