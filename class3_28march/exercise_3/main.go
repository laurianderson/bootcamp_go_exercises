/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle
de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha 
y que tenga un método detalle
*/
package main

import "fmt"

func main() {
	student1 := Student {
		Name:          "Jhon",
		LastName:      "Anderson",
		DNI:           "40.126.568",
		AdmissionDate: "28/03/2023",
	}
	student1.Detail()
} 

type Student struct {
	Name           string
	LastName       string
	DNI            string
	AdmissionDate  string
}

func (s *Student) Detail() {
	fmt.Printf(" Los datos del estudiante son: \n Name: %s \n LastName: %s \n DNI: %s \n AdmissionDate: %s", s.Name, s.LastName, s.DNI, s.AdmissionDate)
}