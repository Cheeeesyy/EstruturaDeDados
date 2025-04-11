package main

import (
	"fmt"
)

func swap(valor1, valor2 *int) (int, int) {
	temp := *valor1
	*valor1 = *valor2
	*valor2 = temp

	return *valor1, *valor2

}

func main() {

	var a int = 10
	var b int = 20

	fmt.Println("Valores antigos de a e b: ", a, b)

	fmt.Print("Novos valores de a e b: ")
	fmt.Println(swap(&a, &b))

}
