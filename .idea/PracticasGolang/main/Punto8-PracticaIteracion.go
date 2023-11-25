package main

import "fmt"

// Función para calcular el MCM utilizando el MCD
func calcularMCM(a, b int) int {
	mcd := calculatorMCD(a, b)
	return a * b / mcd
}

// Función para calcular el MCD utilizando el algoritmo de Euclides
func calculatorMCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int

	fmt.Print("Ingrese el primer número natural: ")
	fmt.Scan(&num1)

	fmt.Print("Ingrese el segundo número natural: ")
	fmt.Scan(&num2)

	// Validar que ambos números sean naturales
	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Por favor, ingrese números naturales válidos.")
		return
	}

	// Calcular y mostrar el MCM
	mcm := calcularMCM(num1, num2)
	fmt.Printf("El mínimo común múltiplo (MCM) de %d y %d es: %d\n", num1, num2, mcm)
}
