package mypackage

import "fmt"

//CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year int
}

type carPrivate struct{
	brand string
	year int
}

//PrintMessage Imprimir mensaje
func PrintMessage(text string){
	var newCar carPrivate
	newCar.brand = "Ferrari"
	newCar.year = 2020
	fmt.Println(text, newCar)
}
