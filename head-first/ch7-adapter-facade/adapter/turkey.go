package adapter

import "log"

type Turkey interface {
	Gobble()
	Fly()
}

var _ Turkey = (*WildTurkey)(nil)

type WildTurkey struct{}

func NewWildTurkey() *WildTurkey {
	return &WildTurkey{}
}

func (w *WildTurkey) Gobble() {
	log.Println("Gobble gobble")
}

func (w *WildTurkey) Fly() {
	log.Println("I'm flying a short distance")
}
