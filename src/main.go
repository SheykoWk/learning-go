package main

import "fmt"

func main() {
	//declaracion de variables
	helloMessage := "Hello"
	worldMessage := "world"
	//Println
	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %v cursos \n", nombre, cursos)

	//Printf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	println(message)

	//tipo de datos
	fmt.Printf("helloMessage: %T", helloMessage)

}
