package main

import "fmt"

//Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição
//do primeiro elemento igual ao valor no slice. Caso não encontre, retorne -1

func ElementoIgual(slice []int, n1 int) int {

	for i := 0; i < len(slice); i++ {
		if slice[i] == n1 {
			return i
		}
	}

	return -1

}
func main() {
	numeros := []int{1, 5, 3, 6, 4, 8, 2}
	n1 := 7
	resultado := ElementoIgual(numeros, n1)
	fmt.Println(resultado)

}
