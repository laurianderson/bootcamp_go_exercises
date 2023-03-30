/*
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos son:
Legajo
Nombre
DNI
Número de teléfono
Domicilio


Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los datos de un array.
En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.
Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar que todos los datos
a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores. Uno de ellos tendrá
que ser del tipo error para el caso de que se ingrese por parámetro algún valor cero (recordá los valores cero de cada tipo de dato,
ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
“Fin de la ejecución” y “Se detectaron varios errores en tiempo de ejecución”. Utilizá defer para cumplir con este requerimiento.

Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente
para el caso de error retornado).
*/

package main

import (
	"errors"
	"fmt"
)

func main() {
    //Instanciando clientes
    cliente1 := Cliente {
        Legajo: 1,
        Nombre: "Pablo",
        DNI: "12345677",
        NumeroTelefono: "45678",
        Domicilio: "Calle1",
    }
    cliente2 := Cliente {
        Legajo: 3,
        Nombre: "",
        DNI: "12345677",
        NumeroTelefono: "45678",
        Domicilio: "Calle1",
    }
    cliente3 := Cliente {
        Legajo: 2,
        Nombre: "Sandra",
        DNI: "12345678",
        NumeroTelefono: "456898",
        Domicilio: "Calle2",
    }

    //Imprimiendo slide clientes previo a  registrarlos
    fmt.Println("Visualizando base, previo a la registración de clientes: ", Clientes)


    //Registrando los clientes
    registrar(cliente1)
    registrar(cliente2)
    registrar(cliente3)

    //Imprimiendo slide clientes post registrarlos
    fmt.Println(Clientes)

    fmt.Println("Fin de ejecución")

}

//Creación slide Clientes y ya tiene un valor de base
var Clientes = []Cliente{{Legajo: 1,
    Nombre: "Pablo",
    DNI: "12345677",
    NumeroTelefono: "45678",
    Domicilio: "Calle1",}}



type Cliente struct {
    Legajo                 int
    Nombre                 string
    DNI                    string
    NumeroTelefono         string
    Domicilio              string
}


func validarExistencia(clienteIngresado Cliente) (bool, error) {
    //Guardo el error del panic, para poder seguir con la ejecución del programa una vez que entra ahí
    defer func(){
        err := recover()
        if err!= nil {
            fmt.Println(err)
        }
    }()
    
    //Recorro el slice y hago la comparación si son iguales primero arrojo true para que el valor cambie y luego por fuera arrojo el panic
    for _, cliente := range Clientes {
        if clienteIngresado.Legajo ==  cliente.Legajo {
            return true, nil
        }
        panic("Error: el cliente con nombre: "  + cliente.Nombre + " y DNI: " + cliente.DNI + " ya existe")
    }
    return false, nil
}

func validarDatos(cliente Cliente) (int, error) {
    fmt.Println("validando datos del cliente: " , cliente.Nombre)
    if cliente.Legajo == 0 || cliente.Nombre == "" || cliente.DNI == "" || cliente.NumeroTelefono == "" ||   cliente.Domicilio == "" {  
       return 0, errors.New("Error")
    } else {
        return 0, nil
    }
}

func registrar(cliente Cliente) {
    existencia, _ := validarExistencia(cliente)
    if !existencia {
        _, err := validarDatos(cliente)
    if err != nil {
        fmt.Println("hubo un error en la registración del cliente: " , cliente.Nombre)
   } else {
        Clientes = append(Clientes, cliente) 
        }
    }
}



