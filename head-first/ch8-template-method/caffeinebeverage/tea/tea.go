package tea

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch8-template-method/caffeinebeverage/beverage"
)

type Tea struct {
	beverage.Beverage
}

func (t *Tea) Brew() {
	log.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	log.Println("Adding Lemon")
}
