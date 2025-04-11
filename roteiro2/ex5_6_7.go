package main

import (
	"fmt"
)

func main() {
	
	type Produto struct {
		Nome       string
		Preco      float64
		Quantidade int
	}

	p := Produto{Nome: "banana", Preco: 3.54, Quantidade: 24}

	fmt.Printf("Produto: %s, Preço: %.2f, Quantidade: %d", p.Nome, p.Preco, p.Quantidade)	
	
}
