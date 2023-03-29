/*
Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba por parámetro 
el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “Error: 
el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int
pasado por parámetro).
*/

package main

import "fmt"

func main() {
	salary := 9000

}

err := fmt.Errorf("Error: el mínimo imponible es de 150.000 y el %s ingresado es de: ", salary)

func salaryIs (salary int , err string) {
	if salary <= 10000 {
		return err
	} else {
		return 0
	}
}