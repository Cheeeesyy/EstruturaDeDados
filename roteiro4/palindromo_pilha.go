package main

import (
	"fmt"
)

// Definição da pilha
type Stack struct {
	items []rune
}

// Push adiciona um caractere na pilha
func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}

// Pop remove e retorna o elemento do topo da pilha
func (s *Stack) Pop() (rune, bool) {
	if len(s.items) == 0 {
		return 0, false // Retorna falso se a pilha estiver vazia
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex] // Remove o último elemento
	return item, true
}

func isPalindrome(word string) bool {

	items := []rune(word)

	stack := Stack{}

	for _, item := range items {

		stack.Push(item)

	}

	for _, item := range items {

		poppedItem, _ := stack.Pop()

		if poppedItem != item {

			return false

		}

	}

	return true

}
func main() {
	// Testes
	words := []string{"radar", "arara", "golang", "level", "hello"}
	for _, word := range words {
		if isPalindrome(word) {
			fmt.Printf("\"%s\" é um palíndromo\n", word)
		} else {
			fmt.Printf("\"%s\" NÃO é um palíndromo\n", word)
		}
	}
}
