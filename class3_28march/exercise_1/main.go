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
	product2 := Product {2, "lemon", 50.50, "lemon fruit", "fruit"}

	fmt.Println("Guardando producto: ")
	product2.Save()

	fmt.Println("Listando producto: ")
	product2.GetAll()

	fmt.Println("Encontrando producto por id: ")
	fmt.Println(getById(1))


}

//Global product instantiation
var Products = []Product {{1, "cake", 150.50, "carrot cake", "cake shop"}}

//Structure
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}


//Methods
func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, p := range Products {
		fmt.Println(p)	
	}
}

//Function
func getById(id int) Product {
	for _, p := range Products {
		if p.ID == id {
			return p
		}else {
			fmt.Println("The product with id " , p.ID , " does not match the id entered")
		}
	}
	return Product{}
}


