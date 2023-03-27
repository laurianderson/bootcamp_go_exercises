/*Testeamos la función calcular impuesto declarado en el archivo main.go*/
package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {
	//Arrange
	sueldo := 51000.00
	descuentoEsperado := 8670.00

	//Act
	descuento := impuestoSalario(sueldo)

	//Assert
	assert.Equal(t, descuentoEsperado, descuento, "Error en el test")
}

