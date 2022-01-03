package main

import "fmt"

func main() {
	//Const declaration
	const pi float64 = 3.1416
	const pi2 = 3.141564
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)
	//Declaracion de variables enteras

	base := 10

	var altura int = 14
	area := base * altura
	fmt.Println("area:", area)

	//ZERO values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

}
