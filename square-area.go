package main

import "fmt"

func main() {
	var side float64

	fmt.Println("******* Calcular el área del cuadrado *******")

	fmt.Print("- Ingrese el lado del cuadrado: ")
	fmt.Scanf("%f", &side)

	res := side * side
	fmt.Println("- El resultado es: ", res)
}