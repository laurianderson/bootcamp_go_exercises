/* 
Crear un programa que cumpla los siguiente puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.   
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de 
Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los 
productos guardados en el slice Products. Una función getById() al cual se le deberá pasar un INT como parámetro y 
retorna el producto correspondiente al parámetro pasado. 
Ejecutar al menos una vez cada método y función definido desde main().
*/

package main

import "fmt"

func main() {
	//Global product instantiation
	Products := []Product {1, "Cake", 150.50, "carrot cake", "cake shop"}

}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}


//Methods
func (p *Product) Save() {
	p.save(Produdcts)

}

func (p *Product) GetAll() {
	fmt.Printf()

}

//Function
func getById(id int) producto {

}


