package main

import (
	"errors"
	"fmt"
	"sort"
)

//Escreva uma função que receba um slice de inteiros como parâmetro
//e retorne um novo slice com os valores ordenados de
//forma crescente. Caso o slice esteja vazio, retorne um erro.

func ordemCrescente(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("Sua slice esta vazia.")
	}
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice, nil

}

func main() {
	slice := []int{5, 2, 7, 1, 3}
	novaSlice, err := ordemCrescente(slice)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(novaSlice)
}
