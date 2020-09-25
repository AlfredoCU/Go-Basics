package main

// Importando nuestras bibliotecas.
import "fmt"

// Creamos una funcion principal.
func main() {
	// Declaramos variables de tipo float64.
	var (
		base   float64
		height float64
	)

	fmt.Println("******* Calcular el área del triángulo *******")

	fmt.Print("- Ingrese la base: ")
	fmt.Scan(&base) // Tomamos de entrada un valor numérico.
	fmt.Print("- Ingrese la altura: ")
	fmt.Scan(&height) // Tomamos de entrada un valor numérico.

	// Formula para sacar el área de un triángulo.
	res := (base * height) / 2
	fmt.Println("- El resultado es: ", res)
}
