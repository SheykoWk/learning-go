package main

import "fmt"


func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	//Recorrer un map

	for i, value:= range m{
		fmt.Println(i, value)
	}

	//Encontrar un valor
	value, ok := m["jose"]
	fmt.Println(value, ok)

}
