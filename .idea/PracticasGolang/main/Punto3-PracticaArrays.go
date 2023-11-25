package main

import (
	"fmt"
)

func transformarCadena(cadena string) string {

	resultadoCadena := ""

	for _, letra := range cadena {

		switch letra {
		case 'a':
			resultadoCadena += "4"
		case 'e':
			resultadoCadena += "3"
		case 'i':
			resultadoCadena += "1"
		case 'o':
			resultadoCadena += "0"
		case 'u':
			resultadoCadena += "6"
		default:
			resultadoCadena += string(letra)
		}
	}
	return resultadoCadena
}

func main() {

	var inputCadena string
	fmt.Printf("Ingrese una palabra cualquiera: ")
	fmt.Scan(&inputCadena)

	resultado := transformarCadena(inputCadena)
	fmt.Println("cadena transformada", resultado)
}
