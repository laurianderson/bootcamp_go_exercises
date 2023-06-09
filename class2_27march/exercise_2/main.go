/*
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y 
devuelva el promedio. No se pueden introducir notas negativas.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(calcularPromedioNotas(-1, 7, 10))
}

func calcularPromedioNotas(notas ...float64) float64 {
	var suma float64

	for _, items := range notas {
		if items >= 0 {
			suma = suma + items 
		} else {
			fmt.Println("No acepta notas negativas")
			break
		}
	}
	return suma / float64(len(notas))
}