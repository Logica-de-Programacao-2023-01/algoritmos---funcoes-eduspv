package main

import (
	"errors"
	"fmt"
	"sort"
)

// Crie uma função que receba um slice de strings como parâmetro e retorne
// uma nova string com as strings em ordem alfabética. Caso o slice esteja vazio, retorne um erro.

func OrdemAlfabetica(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("a sua string esta vazia")
	}
	slice := []rune(str)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	stringOrganizada := string(slice)

	return stringOrganizada, nil
}
func main() {
	string := "eu te amo"
	resultado, err := OrdemAlfabetica(string)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(resultado)
}
