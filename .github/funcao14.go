package main

import (
	"errors"
	"fmt"
)

// Escreva uma função que receba dois slices de inteiros como parâmetros e retorne um novo slice
// contendo apenas os números presentes nos dois slices. Caso um dos slices esteja vazio, retorne um erro.
func numeros(slc1 []int, slc2 []int) ([]int, error) {
	novaslice := []int{}
	if len(slc1) == 0 || len(slc2) == 0 {
		return nil, errors.New("uma das slices está vazia")
	}
	for _, num1 := range slc1 {
		for _, num2 := range slc2 {
			if num1 == num2 {
				novaslice = append(novaslice, num1)
				break
			}
		}
	}

	if len(novaslice) == 0 {
		return nil, errors.New("nenhum número em comum encontrado")
	}
	return novaslice, nil
}
func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{1, 3, 7, 8}
	novaSlice, err := numeros(slice1, slice2)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(novaSlice)

}
