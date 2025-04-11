package main

import (
	"fmt"
)

func main() {
	
	// indice, ra, nome do aluno
	var matriz [2][3]string
	var indice string
	var ra string
	var nomeAluno string

	for i:= 0; i < len(matriz); i++{

		fmt.Println("Insira o índice")
		fmt.Scan(&indice)
		matriz[i][0] = indice

		fmt.Println("Insira o RA")
		fmt.Scan(&ra)
		matriz[i][1] = string(ra)

		fmt.Println("Insira o Nome do Aluno")
		fmt.Scan(&nomeAluno)
		matriz[i][2] = nomeAluno

	}

	fmt.Println("Dados dos alunos:")
	fmt.Println("-----------------")
	fmt.Println("Indice\tRa\tNome\t")

	for i := 0; i < len(matriz); i++{

		fmt.Println(matriz[i][0], matriz[i][1], matriz[i][2])

	}
}