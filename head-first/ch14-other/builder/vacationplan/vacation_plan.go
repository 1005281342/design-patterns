package vacationplan

import (
	"time"
)

type Builder interface {
	BuildDay(date time.Time)
	AddHotel(date time.Time, name string)
	AddTickets(tickets ...string)
	AddReservation(reserve bool)
	AddSpecialEvent(events ...string)
	GetVacationPlan() *VacationPlan
}

// VacationPlan 结构体表示度假规划的详细信息
type VacationPlan struct {
	days        map[time.Time]*DayPlan
	reservation bool
	events      []string
}

func (v *VacationPlan) GetDays() map[time.Time]*DayPlan {
	return v.days
}

func (v *VacationPlan) GetReservation() bool {
	return v.reservation
}

func (v *VacationPlan) GetEvents() []string {
	return v.events
}

// DayPlan 结构体表示每天的规划信息
type DayPlan struct {
	hotel   string
	tickets []string
}

func (d *DayPlan) GetHotel() string {
	return d.hotel
}

func (d *DayPlan) GetTickets() []string {
	return d.tickets
}

// VacationBuilder 实现了 Builder 接口，用于构建度假规划
type VacationBuilder struct {
	vacationPlan *VacationPlan
	currentDate  time.Time
}

func NewVacationBuilder() *VacationBuilder {
	return &VacationBuilder{
		vacationPlan: &VacationPlan{
			days: make(map[time.Time]*DayPlan),
		},
	}
}

func (v *VacationBuilder) BuildDay(date time.Time) {
	v.currentDate = date
	v.vacationPlan.days[date] = &DayPlan{}
}

func (v *VacationBuilder) AddHotel(date time.Time, name string) {
	day := v.vacationPlan.days[date]
	if day != nil {
		day.hotel = name
	}
}

func (v *VacationBuilder) AddTickets(tickets ...string) {
	day := v.vacationPlan.days[v.currentDate]
	if day != nil {
		day.tickets = append(day.tickets, tickets...)
	}
}

func (v *VacationBuilder) AddReservation(reserve bool) {
	v.vacationPlan.reservation = reserve
}

func (v *VacationBuilder) AddSpecialEvent(events ...string) {
	v.vacationPlan.events = append(v.vacationPlan.events, events...)
}

func (v *VacationBuilder) GetVacationPlan() *VacationPlan {
	return v.vacationPlan
}
