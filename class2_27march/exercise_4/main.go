/*
Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas. Por el momento solo tienen tarántulas, hamsters, perros y gatos, 
pero se espera que puedan darle refugio a muchos animales más.
Tienen los siguientes datos:
Perro 10 kg de alimento.
Gato 5 kg de alimento.
Hamster 250 g de alimento.
Tarántula 150 g de alimento.
Se solicita:
-Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y que retorne una función y 
un mensaje (en caso que no exista el animal).
-Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
*/

package main

import "fmt"

func main() {
	animalType := "hamster"
	animal, err := calcularAlimento(animalType)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("La cantidad de alimento para %s es de %f kg \n", animalType, float64(animal(10)) / float64(1000))
	}

}

// Calculamos la cantidad de alimento según el animal
func alimentoPerro(cantidad int) int {
	return cantidad * 10000
}

func alimentoGato(cantidad int) int {
	return cantidad * 5000
}

func alimentoHamster(cantidad int) int {
	return cantidad * 250
}

func alimentoTarantula(cantidad int) int {
	return cantidad * 150
}

func calcularAlimento(tipoAnimal string) (func(cantidadAlimento int) int, error) {
	switch tipoAnimal {
	case "perro":
		return alimentoPerro, nil
	case "gato":
		return alimentoGato, nil
	case "hamster":
		return alimentoHamster, nil
	case "tarantula":
		return alimentoTarantula, nil
	}
	return nil, fmt.Errorf("El animal ingresado no existe en nuestro sistema")
}

