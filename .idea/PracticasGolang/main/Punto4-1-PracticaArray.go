package main

import "fmt"

type secuenciaADN struct {
	Matriz [][]rune
}

func newSecuenciaADN(matriz [][]rune) *secuenciaADN {
	return &secuenciaADN{Matriz: matriz}
}

func (s *secuenciaADN) coincide(s2 *secuenciaADN) bool {

	if len(s.Matriz) != len(s2.Matriz) {
		return false
	}

	for i := 0; i < len(s.Matriz); i++ {
		for j := 0; j < len(s.Matriz[0]); j++ {
			if s.Matriz[i][j] != s2.Matriz[i][j] {
				return false
			}
		}
	}
	return true
}

func (*secuenciaADN) saludar() {
	fmt.Println("hola")
}

func main() {
	matriz1 := [][]rune{
		{'A', 'T', 'G', 'C', 'G', 'T', 'A', 'T'},
		{'T', 'A', 'A', 'T', 'G', 'C', 'G', 'T'},
	}

	matriz2 := [][]rune{
		{'A', 'T', 'A', 'T', 'G', 'C', 'G', 'T'},
		{'A', 'T', 'G', 'C', 'G', 'T', 'A', 'T'},
	}

	secuencia1 := newSecuenciaADN(matriz1)
	secuencia2 := newSecuenciaADN(matriz2)

	if secuencia1.coincide(secuencia2) {
		fmt.Println("las secuencias coinciden")
	} else {
		fmt.Println("las secuencias no coindiden")
	}

	secuencia1.saludar()
}
