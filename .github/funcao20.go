package main

import (
	"errors"
	"fmt"
)

//Escreva uma função que receba um slice de strings como parâmetro e retorne um novo slice
//contendo apenas as strings que possuem mais de 5 caracteres. Caso o slice esteja vazio, retorne um erro.

func stringde5caracteres(slc []string) ([]string, error) {
	NovaSlice := []string{}
	if len(slc) == 0 {
		return nil, errors.New("A slice esta vazia")
	}
	for _, str := range slc {
		caracteres := len(str)
		if caracteres >= 5 {
			NovaSlice = append(NovaSlice, str)
		}
	}
	return NovaSlice, nil
}
func main() {
	slice := []string{"vasco", "da", "gama"}
	resultado, err := stringde5caracteres(slice)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(resultado)
}
