/* Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. 
Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar. 
Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan
más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga no les cobrará interés a 
los que posean un sueldo superior a $100.000. 
Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.*/

package main

import "fmt"

func main(){
	age := 20
	isEmployed := true
	antiquity := 1
	salary := 200.000

	if age < 22 || isEmployed == false || antiquity < 1 {
		fmt.Println("You do not meet the conditions to receive the loan")
		return
	}


	if age > 22  && isEmployed && antiquity > 1 && salary > 100.000 {
		fmt.Println("You can receive the loan")
	} 
	if salary > 100.00 {
		fmt.Println("You can receive the loan without paying interest")
	} 
	if salary < 100.000{
		fmt.Println("You can receive the loan but you should pay interest")
	} 
}