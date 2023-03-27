/* Realizar una aplicación que reciba  una variable con el número del mes. 
Según el número, imprimir el mes que corresponda en texto.*/
package main

import "fmt"

func main() {
	month := make(map[int]string)

	month[1] = "January"
	month[2] = "February"
	month[3] = "March"
	month[4] = "April"
	month[5] = "May"
	month[6] = "June"
	month[7] = "July"
	month[8] = "August"
	month[9] = "September"
	month[10] = "October"
	month[11] = "November"
	month[12] = "December"


	fmt.Println(month[1])
}