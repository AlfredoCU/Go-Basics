package main

// Importando nuestras bibliotecas.
import (
	"fmt"
	"math"
)

// Creamos una funcion principal.
func main() {
	// Declaramos una variable de tipo float64.
	var radio float64

	fmt.Println("******* Calcular el área del círculo *******")

	fmt.Print("- Ingrese el radio del círculo: ")
	fmt.Scanf("%f", &radio) // Tomamos de entrada un valor numérico.

	// Formula para sacar el área de un circulo.
	res := math.Pi * math.Pow(radio, 2)
	fmt.Println("- El resultado es: ", res)
}
