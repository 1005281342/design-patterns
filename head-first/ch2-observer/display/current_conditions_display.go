package display

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch2-observer/observer"
	"github.com/1005281342/design-patterns/head-first/ch2-observer/subject"
)

type CurrentConditionsDisplay struct {
	name string

	temperature float32
	humidity    float32

	subject subject.Subject
}

func NewCurrentConditionsDisplay(name string, subject subject.Subject) (*CurrentConditionsDisplay, error) {
	var display = &CurrentConditionsDisplay{
		name:    name,
		subject: subject,
	}
	if err := subject.Register(name, display); err != nil {
		return nil, err
	}

	return display, nil
}

func (c *CurrentConditionsDisplay) Name() string {
	return c.name
}

func (c *CurrentConditionsDisplay) Display() {
	log.Printf("temperature: %f, humidity:%f\n", c.temperature, c.humidity)
}

func (c *CurrentConditionsDisplay) Update(weather observer.Weather) {
	c.temperature = weather.Temperature
	c.humidity = weather.Humidity
	c.Display()
}
