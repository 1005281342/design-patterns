package beverage

import "log"

type Beverage struct{}

func (b *Beverage) BoilWater() {
	log.Println("Boiling water")
}

func (b *Beverage) Brew() {}

func (b *Beverage) PourInCup() {
	log.Println("Pouring into cup")
}

func (b *Beverage) AddCondiments() {}
