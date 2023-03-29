/*
A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
Ahora que el archivo sí existe, el panic no debe ser lanzado.
Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
En el caso de no poder leerlo, se debe lanzar un “panic”.
Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer” que
nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
*/

package main

import (
	"fmt"
	"io"
	"os"
)


func main() {
	data := readFile("customers.txt")

	fmt.Println(data)

}

func readFile(fileName string) string{
	f1, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	text, err := io.ReadAll(f1)
	if err != nil {
		panic(err)
	}

	defer f1.Close()

	return string(text)
}


	

