package pizza

import (
	"log"

	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizzafactory"
)

type ChicagoStyleCheesePizza struct {
	*pizza
	ingredientFactory *pizzafactory.ChicagoPizzaIngredientFactory
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	return &ChicagoStyleCheesePizza{
		pizza: &pizza{
			name:     "Chicago Style Deep Dish Cheese Pizza",
			dough:    "Extra Thick Crust Dough",
			sauce:    "Plum Tomato Sauce",
			toppings: []string{"Shredded Mozzarella Cheese"},
		},
		ingredientFactory: &pizzafactory.ChicagoPizzaIngredientFactory{},
	}
}

func (c *ChicagoStyleCheesePizza) Cut() {
	log.Println("Cutting the Pizza into square slices")
}

func (c *ChicagoStyleCheesePizza) Prepare() {
	c.pizza.dough = c.ingredientFactory.CreateDough().Name()
	c.pizza.sauce = c.ingredientFactory.CreateSauce().Name()
	c.pizza.cheese = c.ingredientFactory.CreateCheese().Name()
}
