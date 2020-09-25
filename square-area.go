package main

// Importando nuestras bibliotecas.
import "fmt"

// Creamos una funcion principal.
func main() {
	// Declaramos una variable de tipo float64.
	var side float64

	fmt.Println("******* Calcular el área del cuadrado *******")

	fmt.Print("- Ingrese el lado del cuadrado: ")
	fmt.Scanf("%f", &side) // Tomamos de entrada un valor numérico.

	// Formula para sacar el área de un cuadrado.
	res := side * side
	fmt.Println("- El resultado es: ", res)
}
