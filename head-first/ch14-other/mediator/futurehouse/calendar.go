package futurehouse

import "time"

// Calendar 日历
type Calendar struct {
	mediator Mediator

	mockToday *time.Time
}

func (c *Calendar) SetMediator(mediator Mediator) {
	c.mediator = mediator
}

func (c *Calendar) Notify() {}

func (c *Calendar) checkDayOfWeek(date time.Time) bool {
	var weekday = date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func (c *Calendar) CheckDayOfWeek() bool {
	return c.checkDayOfWeek(c.today())
}

func (c *Calendar) MockToday(date time.Time) {
	c.mockToday = &date
}

func (c *Calendar) today() time.Time {
	if c.mockToday != nil {
		return *c.mockToday
	}

	return time.Now()
}

func (c *Calendar) tomorrow() time.Time {
	return c.today().Add(24 * time.Hour)
}

func (c *Calendar) CheckTomorrowOfWeek() bool {
	return c.checkDayOfWeek(c.tomorrow())
}

func (c *Calendar) CheckTomorrowIsGarbageDay() bool {
	var tomorrow = c.tomorrow().Weekday()
	return tomorrow == time.Tuesday || tomorrow == time.Thursday
}
