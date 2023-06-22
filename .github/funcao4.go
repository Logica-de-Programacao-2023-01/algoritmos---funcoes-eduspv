package main

import (
	"fmt"
	"sort"
)

//Crie uma função que receba um slice de inteiros e retorne o segundo maior valor.

func SecondHighest(slice []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(slice)))
	return slice[1]
}

func main() {
	valores := []int{1, 2, 3, 4, 5, 6}
	resultado := SecondHighest(valores)
	fmt.Print(resultado)

}
