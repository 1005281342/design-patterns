package coffee

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch8-template-method/caffeinebeverage/beverage"
)

type Coffee struct {
	beverage.Beverage
}

func (c *Coffee) Brew() {
	log.Println("Dripping Coffee through filter")
}

func (c *Coffee) AddCondiments() {
	log.Println("Adding Sugar and Milk")
}
