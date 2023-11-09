package display

import (
	"log"
	"math"

	"github.com/1005281342/design-patterns/head-first/ch2-observer/observer"
	"github.com/1005281342/design-patterns/head-first/ch2-observer/subject"
)

type StatisticsDisplay struct {
	name string

	maxTemperature float32
	minTemperature float32
	sumTemperature float32
	count          int

	subject subject.Subject
}

func NewStatisticsDisplay(name string, subject subject.Subject) (*StatisticsDisplay, error) {
	var display = &StatisticsDisplay{
		name: name,

		minTemperature: math.MaxFloat32,

		subject: subject,
	}
	if err := subject.Register(name, display); err != nil {
		return nil, err
	}

	return display, nil
}

func (s *StatisticsDisplay) Name() string {
	return s.name
}

func (s *StatisticsDisplay) Display() {
	log.Printf("avg:%f, min:%f, max:%f", s.sumTemperature/float32(s.count), s.minTemperature, s.maxTemperature)
}

func (s *StatisticsDisplay) Update(weather observer.Weather) {
	s.sumTemperature += weather.Temperature
	s.count++

	if weather.Temperature > s.maxTemperature {
		s.maxTemperature = weather.Temperature
	}
	if weather.Temperature < s.minTemperature {
		s.minTemperature = weather.Temperature
	}

	s.Display()
}
