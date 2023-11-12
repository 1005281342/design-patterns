package pizza

import "github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/pizzafactory"

type NYStyleCheesePizza struct {
	*pizza
	ingredientFactory *pizzafactory.NYPizzaIngredientFactory
}

func NewNYStyleCheesePizza() *NYStyleCheesePizza {
	return &NYStyleCheesePizza{
		pizza: &pizza{
			name:     "NY Style Sauce and Cheese Pizza",
			dough:    "Thin Crust Dough",
			sauce:    "Marinara Sauce",
			toppings: []string{"Grated Reggiano Cheese"},
		},
		ingredientFactory: &pizzafactory.NYPizzaIngredientFactory{},
	}
}

func (n *NYStyleCheesePizza) Prepare() {
	n.pizza.dough = n.ingredientFactory.CreateDough().Name()
	n.pizza.sauce = n.ingredientFactory.CreateSauce().Name()
	n.pizza.cheese = n.ingredientFactory.CreateCheese().Name()
}
