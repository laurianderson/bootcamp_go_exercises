/*
Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones. Para ello, cuentan
con todo el detalle necesario en un archivo .txt.
Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo, no han pasado el
archivo a leer por nuestro programa.
Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt” (recordá lo visto sobre el pkg “os”).
Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar un panic al intentar
leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
*/

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	data := readFile("not-exist.txt")
	fmt.Println(data)
	fmt.Println("Ejecución finalizada")
}

func readFile(fileName string) string{
	//defer and recover guardan el error establecido en panic, para continuar con la ejecución del programa
	defer func(){
		errPanic := recover()
		if errPanic!= nil {
            fmt.Println(errPanic)
        }
	}()


	f1, err := os.Open(fileName)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado " + err.Error())
	}
	
	text, err := io.ReadAll(f1)
	if err != nil {
		panic(err)
	}
	
	return string(text)
}