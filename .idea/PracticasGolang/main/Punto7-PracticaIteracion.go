package main

import "fmt"

// Función para calcular el MCD utilizando el algoritmo de Euclides (analisis numerico para ingenieros)
func calcularMCD(a, b int) int {
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

	// Se valida que ambos números sean naturales
	if num1 <= 0 || num2 <= 0 {
		fmt.Println("Por favor, ingrese números naturales válidos.")
		return
	}

	// Se calcula y se muestra el MCD
	mcd := calcularMCD(num1, num2)
	fmt.Printf("El máximo común divisor (MCD) de %d y %d es: %d\n", num1, num2, mcd)
}
