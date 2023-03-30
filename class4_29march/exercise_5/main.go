/*
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
El mismo tendrá que indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
	persona1,err := calcularSalario(90, 2270)
	if err!= nil {
        fmt.Println(err)
    } else{
		fmt.Println(persona1)
	}

}

func calcularSalario(horasTrabajadas , valorHora float64) (float64, error) {
	if horasTrabajadas < 80.0 {
        return 0, errors.New("Error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	
    }
	if valorHora * horasTrabajadas >= 150000 {
        return valorHora * horasTrabajadas - (valorHora * horasTrabajadas * 0.10), nil
    } else {
		return valorHora * horasTrabajadas, nil
	}
}

    
	
	
	
	
