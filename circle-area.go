package main

// Importamos.
import (
	"fmt"
	"math"
)

// Función main
func main() {
	// variable
	var r float64

	fmt.Println("******* Calcular el área del círculo *******")

	fmt.Print("- Ingrese el radio del círculo: ")
	// entradas
	fmt.Scanf("%f", &r)

	// cambios.
	res := math.Pi * math.Pow(r, 2)
	fmt.Println("- El resultado es: ", res)
}
