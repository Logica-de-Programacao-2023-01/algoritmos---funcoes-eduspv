package main

import (
	"errors"
	"fmt"
	"strings"
)

// Escreva uma função que receba uma string como parâmetro e retorne uma nova string com todas
// as vogais substituídas por "*". Caso a string seja vazia, retorne um erro.
func substituirVogais(str string) (string, error) {
	vogais := "AEIOUaeiou"
	if len(str) == 0 {
		return "", errors.New("A string está vazia")
	}
	for _, v := range vogais {
		str = strings.ReplaceAll(str, string(v), "*")
	}
	return str, nil
}
func main() {
	string := "Vasco da gama"
	resultado, err := substituirVogais(string)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(resultado)
}
