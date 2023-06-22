package main

import (
	"errors"
	"fmt"
)

// Função que recebe um slice de inteiros e uma função como parâmetros,
// aplica a função em cada elemento do slice e retorna a soma de todos os resultados.
// Caso o slice esteja vazio, um erro é retornado.

func aplicarFuncaoESomar(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	soma := 0
	for _, num := range slice {
		resultado := funcao(num)
		soma += resultado
	}

	return soma, nil
}

func quadrado(num int) int {
	return num * num
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	resultado, err := aplicarFuncaoESomar(slice, quadrado)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(resultado)
}
