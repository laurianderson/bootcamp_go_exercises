/*
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar 
el valor del precio total.
La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:
Pequeño: solo tiene el costo del producto
Mediano: el precio del producto + un 3% de mantenerlo en la tienda
Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que 
tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y
los adicionales en caso que los tenga
*/

package main

import (
	"fmt"
	"errors"
)

func main() {
	typeProduct, err := factory(medium, 0.0)
	if err != nil {
		fmt.Println(err)
		} else {
			fmt.Println(typeProduct.CalculatePrice())
		}
		fmt.Println(typeProduct)
	}

type Product interface {
	CalculatePrice() float64
}

//Establezco una constante y uso el tipo de dato iota, para que luego en TypeProduct el campo Type funcione como un enum
const (
	small = iota
	medium 
	big 
)

type TypeProduct struct {
	Type  int
	Price float64
}

//Método de la interfaz
func (t TypeProduct) CalculatePrice() float64 {
	if t.Type == small {
		return t.Price
	}
	if t.Type == medium {
		return t.Price + (t.Price * 0.03)
	}
	if t.Type == big {
		return t.Price + (t.Price * 0.06) + 2500.0
	}
	return 0
}

func factory(typeProduct int, price float64) (Product, error) {
	switch typeProduct {
	case small:
		return TypeProduct{
			Price: price,
		} , nil
	case medium:
		return TypeProduct{
			Price: price,
		} , nil
	case big:
		return TypeProduct{
			Price: price,
		} , nil
	default:
		return nil, errors.New("Not found this product type")
	} 
}







