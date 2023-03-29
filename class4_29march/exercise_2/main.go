/*
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
Creá un error personalizado con un struct que implemente “Error()” con el mensaje “Error: el salario
es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000. La validación debe ser hecha con
a función “Is()” dentro del “main”.
*/

package main

import (
	"fmt"
)

func main() {
	salary := 9000
}


type error interface {
	Error() string
}

type MyError struct {
	Message  string
}

func (e *MyError) Error() string {
	return "Error: el salario es menor a 10.000"
}