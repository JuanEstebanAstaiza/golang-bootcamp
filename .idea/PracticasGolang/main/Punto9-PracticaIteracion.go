package main

import "fmt"

func esCapicua(numero int) bool {
	// Convertir el número a un slice de dígitos
	digitos := []int{}
	n := numero
	for n > 0 {
		digito := n % 10
		digitos = append(digitos, digito)
		n /= 10
	}

	// Comparar los dígitos en posiciones opuestas
	for i, j := 0, len(digitos)-1; i < j; i, j = i+1, j-1 {
		if digitos[i] != digitos[j] {
			return false
		}
	}

	return true
}

func main() {
	var numero int

	fmt.Print("Ingrese un número de 5 cifras: ")
	fmt.Scan(&numero)

	if numero < 10000 || numero > 99999 {
		fmt.Println("Por favor, ingrese un número de 5 cifras válido.")
		return
	}

	if esCapicua(numero) {
		fmt.Println("El número es capicúa.")
	} else {
		fmt.Println("El número no es capicúa.")
	}
}
