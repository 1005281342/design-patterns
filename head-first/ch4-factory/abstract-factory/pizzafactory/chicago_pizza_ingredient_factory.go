package pizzafactory

import (
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/cheese"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/clams"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/dough"
	"github.com/1005281342/design-patterns/head-first/ch4-factory/abstract-factory/sauce"
)

type ChicagoPizzaIngredientFactory struct{}

func (c *ChicagoPizzaIngredientFactory) CreateDough() dough.Dough {
	return dough.NewThickDough()
}

func (c *ChicagoPizzaIngredientFactory) CreateSauce() sauce.Sauce {
	return sauce.NewPlumTomatoSauce()
}

func (c *ChicagoPizzaIngredientFactory) CreateCheese() cheese.Cheese {
	return cheese.NewMozzarellaCheese()
}

func (c *ChicagoPizzaIngredientFactory) CreateClam() clams.Clams {
	return clams.NewFrozenClams()
}
