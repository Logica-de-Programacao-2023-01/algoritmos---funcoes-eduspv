package main

import (
	"errors"
	"fmt"
)

// Crie uma função que receba um número inteiro como
// parâmetro e retorne a soma de todos os seus dígitos.
// Caso o número seja negativo, retorne um erro
func soma(num int) (int, error) {
	soma := 0
	if num < 0 {
		return 0, errors.New("o numero é negativo.")
	}
	for num > 0 {
		digito := num % 10
		soma += digito
		num /= 10
	}
	return soma, nil
}

func main() {
	numero := 123
	resultado, err := soma(numero)
	if err != nil {
		fmt.Println("Erro", err)
		return
	}
	fmt.Println(resultado)

}
