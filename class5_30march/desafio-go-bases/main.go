package main

import (
	"github.com/laurianderson/bootcamp_go_exercises/class1_20march/class5_30march/desafio-go-bases/internal/tickets"
	"fmt"
)
func main() {
	//leer el archivo tickets.csv
	tickets.ReadDataFile("tickets.csv")

	//ejercicio1
	totalTickets,_ := tickets.GetTotalTickets("Brazil")
	fmt.Println(totalTickets)

	//ejercicio2
	totalTickets,_ = tickets.GetMornings("mañana")
	fmt.Println(totalTickets)



}