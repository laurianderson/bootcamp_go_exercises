/*
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”, 
se implemente “errors.New()”.
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
con el mensaje “Error: el salario s menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
*/

package main

import (
	"fmt"
	"errors"
)

var error1 = errors.New("Error: el salario es menor a 10.000")

func main() {
	salary := 10000
	
	if salary <= 10000 {
		fmt.Print(error1)
	} else {
		fmt.Println("El salario es de: ", salary)
	}
}
