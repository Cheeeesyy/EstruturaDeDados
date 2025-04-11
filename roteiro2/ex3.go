package main

import (
	"fmt"
)

func main() {

   codigoRa := [6]int{1,2,4,6,7,5}

	//O i representa o valor, e o v representa a posição
	for i, v := range codigoRa{
		fmt.Printf("Posição %d, valor %d\n", i,v)
	}
}
