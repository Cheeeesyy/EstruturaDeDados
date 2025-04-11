package main

import (
	"fmt"
)

func ClassificarNota(nota int) string {

	switch {
	case nota >= 9:
		return "Excelente!"
	case nota >= 7:
		return "Bom!"
	case nota >= 5:
		return "Regular!"
	case nota >= 3:
		return "Ruim!"
	case nota >= 0:
		return "Muito Ruim!"
	default:
		return ""
	}
}


func main() {

	//VERIFICAR SE O NUMERO É PAR OU IMPAR
	// var numero int
	// fmt.Scan(&numero)
	// fmt.Println(ParOuImpar(numero))

	//CLASSIFICAR NOTA
	var nota int
	fmt.Scan(&nota)
	fmt.Println(ClassificarNota(nota))


}
