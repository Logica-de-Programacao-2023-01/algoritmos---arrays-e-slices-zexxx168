package main

import "fmt"

func main() {
	arr := [10]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	maioresque5 := []float64{}

	for _, num := range arr {
		if num > 5 {
			maioresque5 = append(maioresque5, num)
		}
	}
	fmt.Println(maioresque5)
}
