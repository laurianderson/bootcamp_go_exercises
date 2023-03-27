/*Testeamos la función calcular promedio declarado en el archivo main.go*/
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {
	//Arrange
	notas := []float64 {8, 10, 6} 
	promedioEsperado := 8.0

	//Act
	promedio := calcularPromedioNotas(notas...)

	//Assert
	assert.Equal(t, promedioEsperado, promedio, "Error en el test")
}