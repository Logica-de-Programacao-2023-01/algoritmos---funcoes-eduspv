package main

import "fmt"

// Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.
func concatenacao(str string, str2 string) string {
	somaString := str + str2
	return somaString
}

func main() {
	str := "vasco"
	str2 := " da gama"
	resultado := concatenacao(str, str2)
	fmt.Print(resultado)

}
