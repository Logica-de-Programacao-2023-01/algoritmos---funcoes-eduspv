package main

import (
	"errors"
	"fmt"
	"strings"
)

func ConcatenateStrings(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}

	result := strings.Join(slice, ", ")

	return result, nil
}

func main() {
	strSlice := []string{"Olá", "mundo", "!", "Estou", "concatenando", "strings"}
	result, err := ConcatenateStrings(strSlice)

	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(result)
}
