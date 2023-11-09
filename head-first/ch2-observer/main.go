package main

import (
	"github.com/1005281342/design-patterns/head-first/ch2-observer/display"
	"github.com/1005281342/design-patterns/head-first/ch2-observer/observer"
	"github.com/1005281342/design-patterns/head-first/ch2-observer/subject"
)

func main() {
	var (
		sub        = subject.NewWeather()
		err        error
		current    display.Display
		statistics display.Display
	)
	if current, err = display.NewCurrentConditionsDisplay("current", sub); err != nil {
		panic(err)
	}
	if statistics, err = display.NewStatisticsDisplay("statistics", sub); err != nil {
		panic(err)
	}

	sub.SetMeasurements(observer.Weather{
		Temperature: 80,
		Humidity:    65,
		Pressure:    30.4,
	})

	sub.SetMeasurements(observer.Weather{
		Temperature: 82,
		Humidity:    70,
		Pressure:    29.2,
	})

	sub.Unregister(current.Name())

	sub.SetMeasurements(observer.Weather{
		Temperature: 78,
		Humidity:    90,
		Pressure:    29.2,
	})

	sub.Unregister(statistics.Name())
}
