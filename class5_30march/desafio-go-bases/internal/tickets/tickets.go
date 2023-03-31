package tickets

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

type Ticket struct {
	Id      int
	Name    string
	Email   string
	Country string
	Date    string
	Price   int
}

var Tickets = []Ticket{}

func ReadDataFile(path string) ([]Ticket, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return []Ticket{}, err
	}

	//El texto leido lo parseo a string
	text1 := string(text)
	//El texto lo separo por linea, y que vaya haciendo un salto al final de c/u
	rows := strings.Split(text1, "\n")

	var Tickets []Ticket
	//Voy a recorrer cada línea
	for _, row := range rows {
		if row != "" {
			//E ir separando por campo, tomando como referencia donde esta la ,
			values := strings.Split(row, ",")

			//Convierto de string a int
			id, _ := strconv.Atoi(values[0])
			price, _ := strconv.Atoi(values[0])

			//accedo a cada campo por su indice
			ticket := Ticket{
				Id:          id,
				Name:      values[1],
				Email:     values[2],
				Country:   values[3],
				Date:      values[4],
                Price:      price,
            }   
			//guardo el valor de cada indice en el slide de Tickets
			Tickets = append(Tickets, ticket)  
		}
	}
		return Tickets, nil
}




// ejemplo 1
func GetTotalTickets(destination string) (int error) {}

// ejemplo 2
func GetMornings(time string) (int error) {}

// ejemplo 3
func AverageDestination(destination string, total int) (int error) {}