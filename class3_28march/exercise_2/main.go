/*
Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará
 a gestionar correctamente dichos empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de 
los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar
 el método PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/

package main

import (
	"fmt"
) 

func main() {
	//Instantiation of the structure
	person := Person{1, "Brad Pitt", "28/03/2023"}
	employee := Employee{1, "director", person}

	employee.PrintEmployee()
}

//Structure
type Person struct {
	ID           int
	Name         string
	DateOfBirth  string
}

type Employee struct {
	ID        int
	Position  string
	//Composition
	Person    Person 
}

//Method
func (e Employee) PrintEmployee() {
	fmt.Println(" ID:", e.ID,  "\n Position: ", e.Position, "\n Person: ", e.Person)
}





