package facade

import "log"

type ITheaterLights interface {
	Dim(value int)
	On()
}

type TheaterLights struct{}

func NewTheaterLights() *TheaterLights {
	return &TheaterLights{}
}

func (t *TheaterLights) Dim(value int) {
	log.Printf("Theater Ceiling Lights dimming to %d%s \n", value, "%")
}

func (t *TheaterLights) On() {
	log.Println("Theater Ceiling Lights on")
}
