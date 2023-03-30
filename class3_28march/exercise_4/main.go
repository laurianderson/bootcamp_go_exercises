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


//2 resoluciones: La primera correcta, la segunda menos performante (la guarde, por la implementación del enum)
//...........................................PRIMERA RESOLUCIÓN.......................................................
package main

import (
	"errors"
	"fmt"
)


type Product interface {
	calculateCost() float64
}


//Tipo de producto con la implementación del método calcular precio
type CheapProduct struct {
	Price float64
}

func (smp CheapProduct) calculateCost() float64 {
	return smp.Price
}



type MiddleProduct struct {
	Price float64
}
func (mdp MiddleProduct) calculateCost() float64 {
	return mdp.Price + (mdp.Price * 0.03)
}


type ExpensiveProduct struct {
	Price float64
}
func (exp ExpensiveProduct) calculateCost() float64 {
	return exp.Price + (exp.Price * 0.06) + 2500
}



var (
	expensive = "expensive"
	cheap     = "cheap"
	middle    = "middle"
)

func factory(category string, price float64) (Product, error) {
	switch category {
	case cheap:
		return CheapProduct{
			Price: price,
		}, nil
	case middle:
		return MiddleProduct{
			Price: price,
		}, nil
	case expensive:
		return ExpensiveProduct{
			Price: price,
		}, nil
	default:
		return nil, errors.New("Not found Product")
	}
}


func main() {
	//Resolución primer modo (correcto)
	cheapProduct, err := factory("cheap", 30.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cheapProduct.calculateCost())
	}
	expensiveProduct, err := factory("expensive", 100.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expensiveProduct.calculateCost())
	}
	middleProduct, err := factory("middle", 50.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(middleProduct.calculateCost())
	}


	//Resolución segundo modo(poco performante)
	typeProduct, err := factory1(medium, 100.0)
	if err != nil {
		fmt.Println(err)
		} else {
			fmt.Println(typeProduct.calculateCost())
		}
		fmt.Println(typeProduct)
}


//...........................................SEGUNDA RESOLUCIÓN.......................................................

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
func (t TypeProduct) calculateCost() float64 {
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


func factory1(typeProduct int, price float64) (Product, error) {
	switch typeProduct {
	case small:
		return TypeProduct{
			Type: 0,
			Price: price,
		} , nil
	case medium:
		return TypeProduct{
			Type: 1,
			Price: price,
		} , nil
	case big:
		return TypeProduct{
			Type: 2,
			Price: price,
		} , nil
	default:
		return nil, errors.New("Not found this product type")
	} 
}







