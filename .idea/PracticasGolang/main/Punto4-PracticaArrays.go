package main

import "fmt"

type SecuenciaADN struct {
	Cadena string
}

func NewSecuenciaADN(cadena string) *SecuenciaADN {
	return &SecuenciaADN{Cadena: cadena}
}

func (s *SecuenciaADN) Coincide(otra *SecuenciaADN) bool {
	if len(s.Cadena) != len(otra.Cadena) {
		return false
	}

	// Verificar la coincidencia para cada posible punto de inicio
	for i := 0; i < len(s.Cadena); i++ {
		coincide := true
		for j := 0; j < len(s.Cadena); j++ {
			if s.Cadena[(i+j)%len(s.Cadena)] != otra.Cadena[j] {
				coincide = false
				break
			}
		}
		if coincide {
			return true
		}
	}

	return false
}

func main() {
	seq1 := NewSecuenciaADN("ATGCGTAT")
	seq2 := NewSecuenciaADN("ATATGCGT")

	if seq1.Coincide(seq2) {
		fmt.Println("Las secuencias de ADN coinciden.")
	} else {
		fmt.Println("Las secuencias de ADN no coinciden.")
	}
}
