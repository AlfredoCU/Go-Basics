package main

// Importamos nuestra biblioteca.
import "fmt"

// Creamos una funcion principal.
func main() {
	// Declaramos una variable de tipo float64.
	var fahrenheit float64

	fmt.Println("******* Convertir grados Fahrenheit a Celsius *******")

	fmt.Print("- Ingrese los grados ºF: ")
	fmt.Scan(&fahrenheit) // Tomamos de entrada un valor numérico.

	// Formula para convertir grados ºF a ºC.
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println("- Los grados a ºC: ", celsius)
}
