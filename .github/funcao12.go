// Escreva uma função que receba um slice de strings como
// parâmetro e retorne uma string contendo apenas as strings
// que começam com uma letra maiúscula. Caso o slice esteja vazio,
// retorne um erro.
package main

import (
	"errors"
	"strings"
)

func LetraMaiuscula(strs []string) (string, error) {
	if len(strs) == 0 {
		return "", errors.New("Slice está vazio")
	}

	var novaString strings.Builder

	for _, str := range strs {
		if len(str) > 0 && strings.ToUpper(string(str[0])) == string(str[0]) {
			novaString.WriteString(str)
			novaString.WriteString(" ")
		}
	}

	return strings.TrimSpace(novaString.String()), nil
}

func main() {
	strs := []string{"Erro", "lideres", "porfavor", "vasco"}
	resultado, err := LetraMaiuscula(strs)
	if err != nil {
		panic(err)
	}
	println(resultado)
}
