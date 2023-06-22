//Crie uma função que receba um número inteiro como parâmetro
//e retorne verdadeiro se ele for um número primo e falso caso contrário.
//Caso o número seja negativo, retorne um erro

package main

import (
	"fmt"
	"math"
)

func isPrime(num int) (bool, error) {
	if num < 0 {
		return false, fmt.Errorf("Número negativo não é válido")
	}

	if num < 2 {
		return false, nil
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func main() {
	num := 17

	prime, err := isPrime(num)
	if err != nil {
		fmt.Println(err)
		return
	}

	if prime {
		fmt.Printf("%d é um número primo\n", num)
	} else {
		fmt.Printf("%d não é um número primo\n", num)
	}
}
