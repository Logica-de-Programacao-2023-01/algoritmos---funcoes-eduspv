package main

import (
	"errors"
	"fmt"
)

// Crie uma função que receba um número inteiro e um slice de inteiros
// como parâmetros e retorne verdadeiro se o número estiver presente no slice e falso caso contrário.
// Caso o slice esteja vazio, retorne um erro.

func verificacao(num int, slc []int) (bool, error) {
	booleano := false
	if len(slc) == 0 {
		return false, errors.New("A slice esta vazia")
	}
	for i := 0; i < len(slc); i++ {
		if slc[i] == num {
			booleano = true
		} else {
			booleano = false
		}
	}
	if booleano == false {
		return false, nil
	}
	return booleano, nil
}

func main() {
	numero := 3
	slice := []int{1, 2, 3, 4}
	resultado, err := verificacao(numero, slice)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(resultado)

}
