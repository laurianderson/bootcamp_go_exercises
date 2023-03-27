/*Testeamos la función calcular salario declarado en el archivo main.go*/
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {
	//Arrange
	minutos := 60 
	categoria := "A"
	impuestoEsperado := 4500.0

	//Act
	impuesto := calcularSalario(minutos, categoria)

	//Assert
	assert.Equal(t, impuestoEsperado, impuesto, "Error en el test")
}