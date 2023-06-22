package main

import (
	"errors"
	"fmt"
)

// Função que recebe um slice de inteiros e uma função como parâmetros
// A função aplica a função recebida em cada elemento do slice e retorna um novo slice com os resultados
// Caso o slice esteja vazio, um erro é retornado

func ApplyFunctionToSlice(slice []int, fn func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	result := make([]int, len(slice))

	for i, value := range slice {
		result[i] = fn(value)
	}

	return result, nil
}
func main() {
	intSlice := []int{1, 2, 3, 4, 5}
	doubleFn := func(n int) int {
		return 2 * n
	}
	result, err := ApplyFunctionToSlice(intSlice, doubleFn)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(result) // Saída: [2 4 6 8 10]
}
