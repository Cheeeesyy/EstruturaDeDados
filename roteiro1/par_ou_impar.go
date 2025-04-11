package main

import (
	"fmt"
)

func ParOuImpar(num int) string {
	if num%2 == 0 {
		return "Par!"
	} else {
		return "Impar!"
	}
}


func main() {

	//VERIFICAR SE O NUMERO É PAR OU IMPAR
	var numero int
	fmt.Scan(&numero)
	fmt.Println(ParOuImpar(numero))

}
