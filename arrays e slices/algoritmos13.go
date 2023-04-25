package main

import "fmt"

func main() {
	array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var num int

	fmt.Println("digite o numero que será adicionado ao primeiro e ultimo elemento")
	fmt.Scan(&num)

	array[0] = num
	array[6] = num

	fmt.Println(array)

}
