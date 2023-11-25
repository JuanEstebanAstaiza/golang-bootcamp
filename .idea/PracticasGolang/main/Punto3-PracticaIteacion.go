package main

import "fmt"

func main() {
	var number int

	fmt.Printf("ingrese 1 para sumar los 3 primeros #ros naturales y 2 para los primeros 10 #ros naturales: %d")
	fmt.Scan(&number)

	sumaNumeros(number)

}

func sumaNumeros(num int) int {

	suma := 0
	if num == 1 {
		for i := 1; i <= 3; i++ {
			suma += i
		}
		fmt.Println("El resultado de la suma fue:", suma)
	}
	if num == 2 {

		for x := 1; x <= 10; x++ {
			suma += x
		}
		fmt.Println("El resultado de la suma fue:", suma)

	}
	return suma

}
