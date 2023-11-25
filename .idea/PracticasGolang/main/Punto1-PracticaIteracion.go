package main

import (
	"fmt"
	"math"
)

func encontrarMayor(nums []int) int {
	mayor := math.MinInt64

	for _, num := range nums {
		if num > mayor {
			mayor = num
		}
	}

	return mayor
}

func main() {
	var numeros []int
	var cambiapieles int

	fmt.Print("Digite 1 para ingresar 3 numeros y 2 para ingresar 10 numeros: ")
	fmt.Scan(&cambiapieles)

	if cambiapieles == 2 {
		for i := 1; i <= 10; i++ {
			var num int
			fmt.Printf("Ingrese el número %d: ", i)
			fmt.Scan(&num)
			numeros = append(numeros, num)
		}

		resultado := encontrarMayor(numeros)
		fmt.Println("El mayor de los números ingresados es:", resultado)
	}

	if cambiapieles == 1 {
		for u := 1; u <= 3; u++ {
			var nums int
			fmt.Printf("Ingrese el número %d: ", u)
			fmt.Scan(&nums)
			numeros = append(numeros, nums)
		}

		resultados := encontrarMayor(numeros)
		fmt.Println("El mayor de los números ingresados es:", resultados)
	}
}
