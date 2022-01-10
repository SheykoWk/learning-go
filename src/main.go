package main

import (
	"fmt"
	pk "go-basico/src/mypackage"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ford"
	myCar.Year = 1990
	fmt.Println(myCar)

	pk.PrintMessage("holis")
}
