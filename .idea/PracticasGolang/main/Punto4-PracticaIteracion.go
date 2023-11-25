package main

import "fmt"

func sumarPrimerosNaturales(n int) int {
	suma := 0

	for i := 1; i <= n; i++ {
		suma += i
	}

	return suma
}

func main() {
	var n int
	fmt.Print("Ingrese el valor de 'n': ")
	fmt.Scan(&n)

	resultado := sumarPrimerosNaturales(n)
	fmt.Printf("La suma de los primeros %d números naturales es: %d\n", n, resultado)
}
