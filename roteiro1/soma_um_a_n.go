package main

import (
	"fmt"
)

func SomaAte(n int) int {

	var soma int = 0

	for i := 0; i <= n; i++ {
		soma += i
	}

	return soma

}

func main() {

	// SOMA DE NUMEROS DE 1 A N
	var num int
	fmt.Scan(&num)
	fmt.Println(SomaAte(num))

}
