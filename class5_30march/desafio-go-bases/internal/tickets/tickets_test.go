package tickets

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//Testeo ejercicio 1
func TestGetTotalTickets(t *testing.T) {
	ReadDataFile("../../tickets.csv")
		//Arrange
		country := "Brazil"
		expect := 45
		//Act
		totalResult,_ := GetTotalTickets(country)
		//Assertion
		assert.Equal(t, expect, totalResult)
}

//Testeo ejercicio 2
func TestGetTotalMornings(t *testing.T){
	//Arrange
	time := "mañana"
	expect := 256
	//Act
	totalResult,_ := GetMornings(time)
	//Assertion
	assert.Equal(t, expect, totalResult)
}

