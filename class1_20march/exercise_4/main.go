/* Hacer una función que cuente la cantidad total de letras que tiene una palabra y mostrar la letra según el lugar que ocupe en el indice */
package main

import "fmt"

func main() {
	word := []string{"h", "e", "l", "l", "o"} 
	counterLetter := 0

	for i, letter := range word {
		counterLetter ++
		fmt.Println(i, letter)
	}

	fmt.Println("Cantidad de letras:", counterLetter)
}