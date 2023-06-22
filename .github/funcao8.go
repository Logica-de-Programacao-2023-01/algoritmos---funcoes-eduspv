package main

import "fmt"

// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice
// com apenas os números pares contidos no slice. Caso o slice esteja vazio, retorne um erro.
func numerosPares(slice []int) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("A sua slice esta vazia.")
	}
	NovaSlice := []int{}
	for _, pares := range slice {
		if pares%2 == 0 {
			NovaSlice = append(NovaSlice, pares)
		}
	}
	return NovaSlice, nil
}
func main() {
	numeros := []int{1, 2, 3, 4, 5}

	resultado, err := numerosPares(numeros)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}
	fmt.Println(resultado)

}
