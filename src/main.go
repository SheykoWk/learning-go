package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2
	if valor1 == 1  {
		fmt.Println("es 1")
	} else {
		fmt.Println("no es 1")
	}
	// With and
	if valor1 == 1 && valor2 == 2{
		fmt.Println("Es verdad")
	} 

	//with or
	if valor1 == 0 ||  valor2 == 2{
		fmt.Println("Alguna de las 2 es correcta")
	}

}

