package main

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch8-template-method/caffeinebeverage"
	"github.com/1005281342/design-patterns/head-first/ch8-template-method/caffeinebeverage/coffee"
	"github.com/1005281342/design-patterns/head-first/ch8-template-method/caffeinebeverage/tea"
)

func main() {
	log.Println("Making tea")
	var teaBeverage = caffeinebeverage.New(&tea.Tea{})
	teaBeverage.PrepareRecipe()

	log.Println("Making coffee")
	var coffeeBeverage = caffeinebeverage.New(&coffee.Coffee{})
	coffeeBeverage.PrepareRecipe()
}
