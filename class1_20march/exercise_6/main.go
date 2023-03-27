/*Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. 
	 1.Saber cuántos de sus empleados son mayores de 21 años.
	 2.Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	 3.Eliminar a Pedro del mapa.*/
package main

import "fmt"

func main() {
	employees := map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}


	counterPerson := 0
	//1. Employees age > 25
	for person, age := range employees {
		if age > 21 {
			 counterPerson ++
			 fmt.Println(person, age)
		}
	}
	fmt.Println("The number of people over the age of 21 are" , counterPerson)
	 
 
	//2. Add Federico
	fmt.Println("Add Federico")
	employees["Federico"]= 25
	fmt.Println(employees)
 
 
	//3.Delete Pedro
	fmt.Println("Delete Pedro")
	delete(employees, "Pedro")
	fmt.Println(employees)
 
 }