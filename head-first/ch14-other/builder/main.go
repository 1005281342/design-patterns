package main

import (
	"fmt"
	"time"

	"github.com/1005281342/design-patterns/head-first/ch14-other/builder/vacationplan"
)

func main() {
	builder := vacationplan.NewVacationBuilder()

	var today = time.Now()
	builder.BuildDay(today)
	builder.AddHotel(today, "Sample Hotel")
	builder.AddTickets("Ticket 1", "Ticket 2")
	builder.AddReservation(true)
	builder.AddSpecialEvent("Concert", "Tour")

	vacationPlan := builder.GetVacationPlan()

	fmt.Println("Vacation Plan Details:")
	for date, day := range vacationPlan.GetDays() {
		fmt.Println("Date:", date.Format("2006-01-02"))
		fmt.Println("Hotel:", day.GetHotel())
		fmt.Println("Tickets:")
		for _, ticket := range day.GetTickets() {
			fmt.Println("-", ticket)
		}
	}
	fmt.Println("Reservation:", vacationPlan.GetReservation())
	fmt.Println("Special Events:")
	for _, event := range vacationPlan.GetEvents() {
		fmt.Println("-", event)
	}
}
