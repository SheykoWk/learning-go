package main

import "fmt"

func idPalindromo(text string) {
	var textReverse string
	fmt.Scan(&textReverse)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}
	if text == textReverse {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
	fmt.Println(textReverse)
}

func main() {
	slice := []string{"hola", "que", "hace"}
	for i, valor := range slice {
		//indice y valor
		fmt.Println(i, valor)
	}
	for _, valor := range slice {
		//solo valor
		fmt.Println(valor)
	}
	for i := range slice {
		//solo indice
		fmt.Println(i)
	}
	idPalindromo("anita lava la tina")
}
