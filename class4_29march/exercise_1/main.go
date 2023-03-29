/*
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario
ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary” sea menor a 150.000. De lo contrario,
tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.
*/

package main

import (
	"fmt"
)

func main() {
	salary := 10000

	if salary <= 150000 {
		error := MyError{}
		fmt.Println(error)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}


type MyError struct {}

func (e MyError) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}





