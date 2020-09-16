package main

import "fmt"

func main()  {
	var fahrenheit float64

	fmt.Println("******* Convertir grados Fahrenheit a Celsius *******")

	fmt.Print("- Ingrese los grados ºF: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5/9
	fmt.Println("- Los grados a ºC: ", celsius)
}