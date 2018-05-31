package goarea

import (
	"math"
)

// Pi se ja sabe.
const Pi = 3.1415

// Circ responsavel area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect area de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
