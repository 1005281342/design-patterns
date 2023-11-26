package futurehouse

import (
	"log"
	"time"
)

// ConcreteMediator 表示具体的中介者对象，负责协调对象之间的交互
type ConcreteMediator struct {
	alarm     *Alarm
	calendar  *Calendar
	coffeePot *CoffeePot
	sprinkler *Sprinkler
}

func NewConcreteMediator(
	alarm *Alarm,
	calendar *Calendar,
	coffeePot *CoffeePot,
	sprinkler *Sprinkler) *ConcreteMediator {
	var mediator = &ConcreteMediator{
		alarm:     alarm,
		calendar:  calendar,
		coffeePot: coffeePot,
		sprinkler: sprinkler,
	}
	mediator.alarm.SetMediator(mediator)
	mediator.calendar.SetMediator(mediator)
	mediator.coffeePot.SetMediator(mediator)
	mediator.sprinkler.SetMediator(mediator)
	return mediator
}

func (c *ConcreteMediator) Mediate(colleague Colleague) {
	switch colleague.(type) {
	case *Alarm:
		log.Printf("日期:%s 闹钟响了\n", c.calendar.today().Format(time.DateOnly))

		if !c.calendar.CheckDayOfWeek() {
			// 不是周末
			c.coffeePot.StartBrewingCoffee()
		}

		if c.calendar.CheckTomorrowOfWeek() {
			// 明天是周末
			c.sprinkler.Reserve()
		}

	case *CoffeePot:
		log.Println("咖啡煮好了")

		// 做些什么事情
	}
}
