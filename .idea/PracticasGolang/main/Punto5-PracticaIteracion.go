package main

import "fmt"

func main() {
	var numero, suma float64

	fmt.Println("Ingrese una serie de números reales. Ingrese 0 para finalizar.")

	for {
		fmt.Print("Ingrese un número: ")
		fmt.Scan(&numero)

		if numero == 0 {
			break // Sale del bucle si se ingresa 0
		}

		suma += numero
	}

	fmt.Printf("La suma de los números ingresados es: %.2f\n", suma)
}
