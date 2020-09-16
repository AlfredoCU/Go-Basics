package main

import "fmt"

func main() {
	var (
		base float64
		height float64
	)

	fmt.Println("******* Calcular el área del triángulo *******")

	fmt.Print("- Ingrese la base: ")
	fmt.Scan(&base)
	fmt.Print("- Ingrese la altura: ")
	fmt.Scan(&height)

	res := (base * height) / 2
	fmt.Println("- El resultado es: ", res)
}