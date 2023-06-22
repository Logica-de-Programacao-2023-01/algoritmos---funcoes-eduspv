package main

import (
	"errors"
	"fmt"
)

//Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice
//contendo todos os números primos menores ou iguais a ele. Caso o número seja negativo, retorne um erro.

func primo(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func NumerosPrimos(num int) ([]int, error) {
	if num < 0 {
		return nil, errors.New("o seu numero é negativo")
	}
	primos := []int{}
	for i := 2; i <= num; i++ {
		if primo(i) {
			primos = append(primos, i)
		}
	}

	return primos, nil
}
func main() {
	numero := 8
	resultado, err := NumerosPrimos(numero)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(resultado)

}
