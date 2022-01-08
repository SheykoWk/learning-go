package main

import "fmt"

func main() {
	
	switch module := 4 % 2; module{
	case 0: 
	fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}
	//switch without condition
	value := 45
	switch{
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}

