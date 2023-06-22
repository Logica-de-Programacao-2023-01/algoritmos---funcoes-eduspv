package main

import (
	"errors"
	"fmt"
	"strings"
)

//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string.
//Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro]

func palavras(str string) ([]string, error) {
	if len(str) == 0 {
		return nil, errors.New("a sua string é vazia")
	}
	palavra := strings.Fields(str)
	return palavra, nil
}
func main() {
	str := "vasco da gama meu bem"
	final, err := palavras(str)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	for _, palavra := range final {
		fmt.Println(palavra)
	}

}
