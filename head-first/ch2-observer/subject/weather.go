package subject

import (
	"fmt"

	"github.com/1005281342/design-patterns/head-first/ch2-observer/observer"
)

type Weather struct {
	observers map[string]observer.Observer
	observer.Weather
}

func NewWeather() *Weather {
	return &Weather{observers: make(map[string]observer.Observer)}
}

func (w *Weather) Register(name string, observer observer.Observer) error {
	if _, hit := w.observers[name]; hit {
		return fmt.Errorf("观察者 %s 已存在", name)
	}

	w.observers[name] = observer
	return nil
}

func (w *Weather) Unregister(name string) {
	delete(w.observers, name)
}

func (w *Weather) Notify() {
	for _, o := range w.observers {
		o.Update(w.Weather)
	}
}

func (w *Weather) MeasurementsChanged() {
	w.Notify()
}

func (w *Weather) SetMeasurements(weather observer.Weather) {
	w.Weather = weather
	w.MeasurementsChanged()
}
