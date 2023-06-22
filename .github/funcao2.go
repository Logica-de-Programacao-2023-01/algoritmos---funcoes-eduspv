package main

import (
	"fmt"
	"strings"
)

//Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

func VogaisPresentes(Str string) int {
	contador := 0
	vogais := "aeiouAEIOU"

	for _, letras := range Str {
		if strings.ContainsRune(vogais, letras) {
			contador++
		}
	}
	return contador

}

func main() {
	Str := "Vasco da gama meu bem."
	vowelCount := VogaisPresentes(Str)
	fmt.Print(vowelCount)

}
