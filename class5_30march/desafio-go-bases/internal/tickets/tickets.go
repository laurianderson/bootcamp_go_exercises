package tickets

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"errors"
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
func GetTotalTickets(destination string) (int, error) {
	var countTicket = 0
	if len(Tickets) <= 0 {
		return 0, errors.New("Not found data")
	}
	for _, ticket := range Tickets {
        if ticket.Country == destination {
			countTicket++ 
        } 
    }
	return countTicket, nil
}

//ejemplo 2
func GetMornings(time string) (int, error) {
	total := 0
	switch time {
	case "madrugada":
		for _,ticket := range Tickets {
			if ticket.Date >= "0" && ticket.Date <= "6" {
				total++
			}
		}
		return total, nil
	case "mañana":
		for _, ticket := range Tickets {
			if ticket.Date >= "7" && ticket.Date <= "12" {
				total++
			}
		}
		return total, nil
	case "tarde":
		for _, ticket := range Tickets {
			if ticket.Date >= "13" && ticket.Date <= "19" {
				total++
			}
		}
		return total, nil
	case "noche":
		for _, ticket := range Tickets {
			if ticket.Date >= "20" && ticket.Date <= "23" {
				total++
			}
		}
		return total, nil
	default:
		return 0, errors.New("Not found time")
	}
}

// ejemplo 3
func AverageDestination(destination string, total int) (int, error) {
	totalCountryTickets, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total <= 0 {
		return 0, errors.New("Canot dive by 0")
	}
	return totalCountryTickets / total, nil
}