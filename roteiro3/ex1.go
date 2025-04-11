package main

import (
	"fmt"
)

func main() {


	var b float64 = 14.44
	var p2 *float64 = &b

	fmt.Println("valor de b antes da troca: ", b)

	*p2 = 18.88

	fmt.Println("valor de b depois da troca: ", b)


}
