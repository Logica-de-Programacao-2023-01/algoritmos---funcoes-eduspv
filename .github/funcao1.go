package main

import "fmt"

// Crie uma função que receba um slice de
// inteiros e retorne a média aritmética dos valores.
func calcularMedia(slice []int) float64 {
	if len(slice) == 0 {
		return 0
	}
	soma := 0
	for _, valor := range slice {
		soma += valor
	}
	return float64(soma) / float64(len(slice))
}

func main() {
	valores := []int{10, 20, 30, 40, 50}
	media := calcularMedia(valores)
	fmt.Println(media)

}
