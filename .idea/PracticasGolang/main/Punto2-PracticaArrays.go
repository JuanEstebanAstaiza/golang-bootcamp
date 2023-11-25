package main

import "fmt"

func borrarElemento(array []int, indice int) []int {
	return append(array[:indice], array[indice+1:]...)
}
func main() {
	var indice int
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Antes de borrar:", slice)
	fmt.Print("Inserte el índice que quiere borrar: ")
	fmt.Scan(&indice)
	// Validar el índice para evitar un índice fuera de rango
	if indice < 0 || indice >= len(slice) {
		fmt.Println("Índice inválido.")
		return
	}
	// Borrar elemento
	slice = borrarElemento(slice, indice)
	fmt.Println("Después de borrar el índice:", slice)
	// Rellenar la posición vacía con cero
	slice = append(slice, 0)
	fmt.Println("Después de modificar:", slice)
}
