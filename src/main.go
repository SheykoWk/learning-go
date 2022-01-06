package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}
func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "hola")
	fmt.Println(returnValue(2))

	value1, value2 := doubleReturn(2)
	value3, _ := doubleReturn(8)

	fmt.Println(value1, value2)
	fmt.Println(value3)
}
