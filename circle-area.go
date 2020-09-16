package main

import (
	"fmt"
	"math"
)

func main()  {
	var radio float64

	fmt.Println("******* Calcular el área del círculo *******")

	fmt.Print("- Ingrese el radio del círculo: ")
	fmt.Scanf("%f", &radio)

	res := math.Pi * math.Pow(radio,2)
	fmt.Println("- El resultado es: ", res)
}