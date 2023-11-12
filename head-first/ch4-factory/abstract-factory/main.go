package main

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizza"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizzastore"
)

func main() {
	var (
		nyStore      = pizzastore.NewNYPizzaStore()
		chicagoStore = pizzastore.NewChicagoPizzaStore()
		piz          pizza.Pizza
		err          error
	)

	if piz, err = nyStore.OrderPizza(pizza.TypeCheese.String()); err != nil {
		log.Fatal(err)
	}
	piz.Info()
	log.Printf("Ethan ordered a %s\n", piz.GetName())

	if piz, err = chicagoStore.OrderPizza(pizza.TypeCheese.String()); err != nil {
		log.Fatal(err)
	}
	piz.Info()
	log.Printf("Joel ordered a %s\n", piz.GetName())
}
